package tests

import (
	"bytes"
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
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type mockStaffService struct {
	CreateFunc func(*entities.Staff, string) error
	LoginFunc  func(string, string, uuid.UUID) (*entities.Staff, error)
}

func (m *mockStaffService) Create(staff *entities.Staff, password string) error {
	return m.CreateFunc(staff, password)
}
func (m *mockStaffService) Login(username, password string, hospitalID uuid.UUID) (*entities.Staff, error) {
	return m.LoginFunc(username, password, hospitalID)
}
func (m *mockStaffService) GetHospitalIDByStaffID(staffID string) (uuid.UUID, error) {
	return uuid.New(), nil
}

func TestStaffHandler_CreateHandler_Positive(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	h := handlers.NewStaffHandler(&mockStaffService{
		CreateFunc: func(staff *entities.Staff, password string) error {
			return nil
		},
	})
	router.POST("/staff", h.CreateHandler)
	body := dto.StaffCreateRequest{
		Username:   "testuser",
		Password:   "password",
		HospitalID: uuid.New(),
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/staff", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestStaffHandler_CreateHandler_Negative_EmptyHospitalID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	h := handlers.NewStaffHandler(&mockStaffService{
		CreateFunc: func(staff *entities.Staff, password string) error {
			return errors.New("fail")
		},
	})
	router.POST("/staff", h.CreateHandler)
	body := dto.StaffCreateRequest{
		Username: "testuser",
		Password: "password",
		//not send HospitalID
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/staff", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusInternalServerError, w.Code, "expected error response")

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	assert.Equal(t, "error", response["status"], "expected error status")
}

func TestStaffHandler_LoginHandler_Positive(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	os.Setenv("JWT_SECRET", "testsecret")
	h := handlers.NewStaffHandler(&mockStaffService{
		LoginFunc: func(username, password string, hospitalID uuid.UUID) (*entities.Staff, error) {
			return &entities.Staff{ID: uuid.New(), Username: username, HospitalID: hospitalID, CreatedAt: time.Now(), UpdatedAt: time.Now()}, nil
		},
	})
	router.POST("/staff/login", h.LoginHandler)
	body := dto.StaffLoginRequest{
		Username:   "testuser",
		Password:   "password",
		HospitalID: uuid.New(),
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/staff/login", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "expected status 200")

}

func TestStaffHandler_LoginHandler_Negative_WrongPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	h := handlers.NewStaffHandler(&mockStaffService{
		LoginFunc: func(username, password string, hospitalID uuid.UUID) (*entities.Staff, error) {
			return nil, errors.New("invalid credentials")
		},
	})
	router.POST("/staff/login", h.LoginHandler)
	body := dto.StaffLoginRequest{
		Username:   "testuser",
		Password:   "wrongpassword",
		HospitalID: uuid.New(),
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/staff/login", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected status 401")
}
