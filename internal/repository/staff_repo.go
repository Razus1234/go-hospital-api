package repository

import (
	"errors"
	"go-hospital-api/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StaffRepository interface {
	Create(staff *entities.Staff) error
	GetByUsername(username string, hospitalID uuid.UUID) (*entities.Staff, error)
	GetHospitalIDByStaffID(staffID string) (uuid.UUID, error)
}
type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) StaffRepository {
	return &staffRepository{db: db}
}

func (r *staffRepository) Create(staff *entities.Staff) error {
	// insert into staff table
	return r.db.Create(staff).Error
}

func (r *staffRepository) GetByUsername(username string, hospitalID uuid.UUID) (*entities.Staff, error) {
	var staff entities.Staff
	if err := r.db.Where("username = ? AND hospital_id = ?", username, hospitalID).First(&staff).Error; err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepository) GetHospitalIDByStaffID(staffID string) (uuid.UUID, error) {
	var staff entities.Staff
	staffUUID, err := uuid.Parse(staffID)
	if err != nil {
		return uuid.Nil, errors.New("invalid staff ID")
	}
	if err := r.db.Where(&entities.Staff{ID: staffUUID}).First(&staff).Error; err != nil {
		return uuid.Nil, errors.New("staff not found")
	}
	return staff.HospitalID, nil

}
