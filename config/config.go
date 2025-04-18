package config

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // Global variable for database connection

// ConnectDatabase initializes the database connection based on DATABASE_URL
func ConnectDatabase() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	if strings.HasPrefix(databaseURL, "postgres://") || strings.HasPrefix(databaseURL, "postgresql://") {
		// Connecting to PostgreSQL database
		DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL database: %v", err)
		}
		log.Println("Connected to PostgreSQL database successfully!")
	} else if strings.HasPrefix(databaseURL, "sqlite3://") {
		// Extract SQLite file path from DATABASE_URL
		sqlitePath := strings.TrimPrefix(databaseURL, "sqlite3://")
		DB, err = gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to SQLite database: %v", err)
		}
		log.Println("Connected to SQLite database successfully!")
	} else {
		log.Fatal("Unsupported database type in DATABASE_URL")
	}
}
