package handlers

import (
	"go-hospital-api/internal/dto"
	"go-hospital-api/internal/services"
	"go-hospital-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientService services.PatientServiceInterface
	staffService   services.StaffServiceInterface
}

func NewPatientHandler(patientService services.PatientServiceInterface, staffService services.StaffServiceInterface) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
		staffService:   staffService,
	}
}

// SearchHandler for patient search with hospital restriction
// @Tags patients
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param search body dto.PatientSearchCriteria  true "Search Criteria"
// @Success 200 {object} []dto.PatientResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /patients/search [post]
func (h *PatientHandler) SearchHandler(c *gin.Context) {
	staffID, err := utils.VerifyTokenFromRequest(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}
	hospitalID, err := h.staffService.GetHospitalIDByStaffID(staffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}
	var req dto.PatientSearchCriteria
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}
	req.HospitalID = hospitalID
	patients, err := h.patientService.Search(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Status: "error", Message: err.Error()})
		return
	}
	resp := make([]dto.PatientResponse, len(patients))
	for i, p := range patients {
		resp[i] = dto.PatientResponse{
			PatientID:    p.ID,
			FirstNameTh:  p.FirstNameTH,
			MiddleNameTh: p.MiddleNameTH,
			LastNameTh:   p.LastNameTH,
			DateOfBirth:  *p.DateOfBirth,
			PatientHN:    p.PatientHN,
			NationalID:   p.NationalID,
			PassportID:   p.PassportID,
			PhoneNumber:  p.PhoneNumber,
			Email:        p.Email,
			Gender:       p.Gender,
			HospitalID:   p.HospitalID,
			CreatedAt:    p.CreatedAt,
			UpdatedAt:    p.UpdatedAt,
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": resp})
}
