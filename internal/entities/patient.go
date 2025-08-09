package entities

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstNameTH  string
	MiddleNameTH string
	LastNameTH   string
	FirstNameEN  string
	MiddleNameEN string
	LastNameEN   string
	DateOfBirth  *time.Time
	PatientHN    string
	NationalID   string
	PassportID   string
	PhoneNumber  string
	Email        string
	Gender       string
	HospitalID   uuid.UUID `gorm:"type:uuid;not null"`
	Hospital     Hospital  `gorm:"foreignKey:HospitalID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
