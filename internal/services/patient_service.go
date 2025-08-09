package services

import (
	"context"
	"go-hospital-api/internal/dto"
	"go-hospital-api/internal/entities"
	"go-hospital-api/internal/repository"
)

type PatientServiceInterface interface {
	Search(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error)
}

type PatientService struct {
	repo repository.PatientRepository
}

func NewPatientService(repo repository.PatientRepository) PatientServiceInterface {
	return &PatientService{repo: repo}
}

// Search returns patients filtered by criteria and hospital restriction
func (s *PatientService) Search(ctx context.Context, criteria dto.PatientSearchCriteria) ([]entities.Patient, error) {
	// Always filter by hospitalID to restrict results
	return s.repo.Search(ctx, criteria)
}
