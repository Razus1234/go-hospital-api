package db

import (
	"go-hospital-api/internal/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGORM(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&entities.Hospital{},
		&entities.Staff{},
		&entities.Patient{},
	); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
