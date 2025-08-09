package dto

import (
	"time"

	"github.com/google/uuid"
)

// ---------------- Error ----------------
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// ---------------- Hospital ----------------
type HospitalRequest struct {
	Name string `json:"name" validate:"required"`
}

type HospitalResponse struct {
	HospitalID uuid.UUID `json:"hospital_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ---------------- Staff ----------------
type StaffCreateRequest struct {
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	HospitalID uuid.UUID `json:"hospital_id" validate:"required"`
}

type StaffLoginRequest struct {
	Username   string    `json:"username" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	HospitalID uuid.UUID `json:"hospital_id" validate:"required"`
}

type StaffLoginResponse struct {
	Token string `json:"token"`
}

type StaffResponse struct {
	ID         uuid.UUID `json:"staff_id"`
	Username   string    `json:"username"`
	HospitalID uuid.UUID `json:"hospital_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ---------------- Patient ----------------
type PatientSearchCriteria struct {
	NationalID  *string    `json:"national_id"`
	PassportID  *string    `json:"passport_id"`
	FirstName   *string    `json:"first_name"`
	MiddleName  *string    `json:"middle_name"`
	LastName    *string    `json:"last_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	PhoneNumber *string    `json:"phone_number"`
	Email       *string    `json:"email"`
	HospitalID  uuid.UUID  `json:"hospital_id"`
}

type PatientResponse struct {
	PatientID    uuid.UUID `json:"patient_id"`
	FirstNameTh  string    `json:"first_name_th"`
	MiddleNameTh string    `json:"middle_name_th"`
	LastNameTh   string    `json:"last_name_th"`
	FirstNameEn  string    `json:"first_name_en"`
	MiddleNameEn string    `json:"middle_name_en"`
	LastNameEn   string    `json:"last_name_en"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	PatientHN    string    `json:"patient_hn"`
	NationalID   string    `json:"national_id"`
	PassportID   string    `json:"passport_id"`
	PhoneNumber  string    `json:"phone_number"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	HospitalID   uuid.UUID `json:"hospital_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
