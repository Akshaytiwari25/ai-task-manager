package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Akshaytiwari25/ai-task-manager-backend/models"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection and migrates models.
func ConnectDatabase() {
	// Load environment variables from .env (if available)
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}

	// Get the DATABASE_URL from environment variables.
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env file")
	}

	// Connect to PostgreSQL using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the models from the models package.
	err = db.AutoMigrate(&models.User{}, &models.Task{})
	if err != nil {
		log.Fatal("Database migration failed:", err)
	}

	DB = db
	fmt.Println("Database connected and migrated successfully!")
}
