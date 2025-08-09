package entities

import (
	"time"

	"github.com/google/uuid"
)

type Staff struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username     string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	HospitalID   uuid.UUID `gorm:"type:uuid;not null"`
	Hospital     Hospital  `gorm:"foreignKey:HospitalID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
