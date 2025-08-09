package main

import (
	"fmt"
	"go-hospital-api/internal/db"
	"go-hospital-api/internal/handlers"
	"go-hospital-api/internal/middleware"
	"go-hospital-api/internal/repository"
	"go-hospital-api/internal/services"
	"log"
	"os"

	_ "go-hospital-api/docs" // Import the docs package to generate Swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Hospital API
// @version 1.0
// @description This is a exam for hospital api.
// @contact.name developer
// @contact.email jakkrid.wi@gmail.com
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT Authorization header using the Bearer scheme. Example: "Authorization: Bearer {token}"
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	dbConn, err := db.ConnectGORM(dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	db.AutoMigrate(dbConn)

	// Wire repositories
	staffRepo := repository.NewStaffRepository(dbConn)
	patientRepo := repository.NewPatientRepository(dbConn)

	// Wire services (use interfaces)
	staffService := services.NewStaffService(staffRepo)
	patientService := services.NewPatientService(patientRepo)

	// Wire handlers (use interfaces)
	staffHandler := handlers.NewStaffHandler(staffService)
	patientHandler := handlers.NewPatientHandler(patientService, staffService)

	r := gin.Default()

	// Swagger UI endpoint (http://localhost:8080/swagger/index.html)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	// Public routes
	api.POST("/staff/create", staffHandler.CreateHandler)
	api.POST("/staff/login", staffHandler.LoginHandler)

	// Authenticated routes
	auth := api.Group("")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/patients/search", patientHandler.SearchHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on :%s", port)
	log.Printf("swagger docs available at http://localhost:%s/swagger/index.html", port)
	r.Run(":" + port)
}
