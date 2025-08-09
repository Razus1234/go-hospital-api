package repository

import (
	"context"
	"go-hospital-api/internal/entities"

	"go-hospital-api/internal/dto"

	"gorm.io/gorm"
)

type PatientRepository interface {
	Search(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error)
}
type patientRepo struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) PatientRepository {
	return &patientRepo{db: db}
}

func (r *patientRepo) Search(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
	var patients []entities.Patient

	query := r.db.WithContext(ctx).Where("hospital_id = ?", criteria.HospitalID)

	if criteria.NationalID != nil && *criteria.NationalID != "" {
		query = query.Where("national_id = ?", *criteria.NationalID)
	}
	if criteria.PassportID != nil && *criteria.PassportID != "" {
		query = query.Where("passport_id = ?", *criteria.PassportID)
	}
	if criteria.FirstName != nil && *criteria.FirstName != "" {
		query = query.Where("first_name_th ILIKE ?", "%"+*criteria.FirstName+"%")
	}
	if criteria.MiddleName != nil && *criteria.MiddleName != "" {
		query = query.Where("middle_name_th ILIKE ?", "%"+*criteria.MiddleName+"%")
	}
	if criteria.LastName != nil && *criteria.LastName != "" {
		query = query.Where("last_name_th ILIKE ?", "%"+*criteria.LastName+"%")
	}
	if criteria.DateOfBirth != nil {
		query = query.Where("date_of_birth = ?", *criteria.DateOfBirth)
	}
	if criteria.PhoneNumber != nil && *criteria.PhoneNumber != "" {
		query = query.Where("phone_number ILIKE ?", "%"+*criteria.PhoneNumber+"%")
	}
	if criteria.Email != nil && *criteria.Email != "" {
		query = query.Where("email ILIKE ?", "%"+*criteria.Email+"%")
	}

	err := query.Find(&patients).Error
	return patients, err
}
