package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"go-hospital-api/internal/dto"
	"go-hospital-api/internal/entities"
	"go-hospital-api/internal/handlers"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockPatientService struct {
	SearchFunc func(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error)
}

func (m *mockPatientService) Search(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
	return m.SearchFunc(ctx, criteria)
}

type mockStaffService2 struct {
	GetHospitalIDByStaffIDFunc func(staffID string) (uuid.UUID, error)
	CreateFunc                 func(staff *entities.Staff, hospitalID string) error
	LoginFunc                  func(email, password string, hospitalID uuid.UUID) (*entities.Staff, error)
}

func (m *mockStaffService2) GetHospitalIDByStaffID(staffID string) (uuid.UUID, error) {
	return m.GetHospitalIDByStaffIDFunc(staffID)
}

func (m *mockStaffService2) Create(staff *entities.Staff, hospitalID string) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(staff, hospitalID)
	}
	return nil
}

func (m *mockStaffService2) Login(email, password string, hospitalID uuid.UUID) (*entities.Staff, error) {
	if m.LoginFunc != nil {
		return m.LoginFunc(email, password, hospitalID)
	}
	return nil, nil
}

func TestPatientHandler_SearchHandler(t *testing.T) {
	cases := []struct {
		name           string
		searchFunc     func(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error)
		getHospitalID  func(staffID string) (uuid.UUID, error)
		token          string
		wantStatusCode int
	}{
		{
			name: "positive",
			searchFunc: func(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
				return []entities.Patient{}, nil
			},
			getHospitalID: func(staffID string) (uuid.UUID, error) {
				return uuid.New(), nil
			},
			token:          generateValidToken(),
			wantStatusCode: http.StatusOK,
		},
		{
			name: "negative search error",
			searchFunc: func(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
				return nil, errors.New("fail")
			},
			getHospitalID: func(staffID string) (uuid.UUID, error) {
				return uuid.New(), nil
			},
			token:          generateValidToken(),
			wantStatusCode: http.StatusInternalServerError,
		},
		{
			name: "negative hospital id error",
			searchFunc: func(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
				return []entities.Patient{}, nil
			},
			getHospitalID: func(staffID string) (uuid.UUID, error) {
				return uuid.Nil, errors.New("fail")
			},
			token:          generateValidToken(),
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			patientService := &mockPatientService{SearchFunc: tc.searchFunc}
			staffService2 := &mockStaffService2{GetHospitalIDByStaffIDFunc: tc.getHospitalID}
			h := handlers.NewPatientHandler(patientService, staffService2)

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.POST("/patients/search", h.SearchHandler)

			criteria := dto.PatientSearchCriteria{}
			b, _ := json.Marshal(criteria)
			req := httptest.NewRequest("POST", "/patients/search", bytes.NewReader(b))
			req.Header.Set("Authorization", tc.token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			assert.Equal(t, tc.wantStatusCode, w.Code)
		})
	}
}

func generateValidToken() string {
	os.Setenv("JWT_SECRET", "testsecret")
	// You'll need to import the same JWT package and secret used in your login handler
	// Example implementation (adjust based on your actual token generation logic):
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uuid.New().String(),
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, _ := token.SignedString([]byte("testsecret"))
	return "Bearer " + tokenString
}
