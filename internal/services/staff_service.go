package services

import (
	"errors"
	"go-hospital-api/internal/entities"
	"go-hospital-api/internal/repository"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

type StaffServiceInterface interface {
	Create(staff *entities.Staff, plainPassword string) error
	Login(username, plainPassword string, hospitalID uuid.UUID) (*entities.Staff, error)
	GetHospitalIDByStaffID(staffID string) (uuid.UUID, error)
}

type StaffService struct {
	repo repository.StaffRepository
}

func NewStaffService(repo repository.StaffRepository) StaffServiceInterface {
	return &StaffService{repo: repo}
}

// Create creates a new staff member with hashed password
func (s *StaffService) Create(staff *entities.Staff, plainPassword string) error {
	if staff.ID == uuid.Nil {
		staff.ID = uuid.New()
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	staff.PasswordHash = string(hashedPwd)
	staff.CreatedAt = time.Now()
	staff.UpdatedAt = time.Now()

	return s.repo.Create(staff)
}

// Login verifies username and password, returns staff if successful
func (s *StaffService) Login(username, plainPassword string, hospitalID uuid.UUID) (*entities.Staff, error) {
	staff, err := s.repo.GetByUsername(username, hospitalID)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	// compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(staff.PasswordHash), []byte(plainPassword))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	return staff, nil
}

func (s *StaffService) GetHospitalIDByStaffID(staffID string) (uuid.UUID, error) {
	hospitalID, err := s.repo.GetHospitalIDByStaffID(staffID)
	if err != nil {
		return uuid.Nil, err
	}

	return hospitalID, nil
}
