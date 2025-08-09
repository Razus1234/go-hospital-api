package handlers

import (
	"go-hospital-api/internal/dto"
	"go-hospital-api/internal/entities"
	"go-hospital-api/internal/services"
	"go-hospital-api/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
)

type StaffHandler struct {
	service services.StaffServiceInterface
}

func NewStaffHandler(service services.StaffServiceInterface) *StaffHandler {
	return &StaffHandler{service: service}
}

// CreateHandler creates new staff
// @Summary Create staff
// @Description Create a new staff
// @Tags staff
// @Accept json
// @Produce json
// @Param staff body dto.StaffCreateRequest true "Staff info"
// @Success 200 {object} dto.StaffResponse
// @Failure 400 {object} dto.ErrorResponse "bad request"
// @Failure 500 {object} dto.ErrorResponse "internal server error"
// @Router /staff/create [post]
func (h *StaffHandler) CreateHandler(c *gin.Context) {
	var req dto.StaffCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	staff := entities.Staff{
		ID:         uuid.New(),
		Username:   req.Username,
		HospitalID: req.HospitalID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	if err := h.service.Create(&staff, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	resp := dto.StaffResponse{
		ID:         staff.ID,
		Username:   staff.Username,
		HospitalID: staff.HospitalID,
		CreatedAt:  staff.CreatedAt,
		UpdatedAt:  staff.UpdatedAt,
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": resp})
}

// LoginHandler authenticates staff and returns JWT token
// @Summary Login staff
// @Description Login an existing staff
// @Tags staff
// @Accept json
// @Produce json
// @Param staff body dto.StaffLoginRequest true "Staff login info"
// @Success 200 {object} dto.StaffLoginResponse
// @Failure 400 {object} dto.ErrorResponse "bad request"
// @Failure 401 {object} dto.ErrorResponse "unauthorized"
// @Failure 500 {object} dto.ErrorResponse "internal server error"
// @Router /staff/login [post]
func (h *StaffHandler) LoginHandler(c *gin.Context) {
	var req dto.StaffLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}

	staff, err := h.service.Login(req.Username, req.Password, req.HospitalID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Status: "error", Message: "invalid username or password"})
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Status: "error", Message: "JWT_SECRET not configured"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(staff.ID.String(), jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Status: "error", Message: "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": dto.StaffLoginResponse{Token: token}})
}
