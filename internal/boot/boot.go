package boot

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/garvit4540/simplepay/internal/database/migrations"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// Initialize initializes the application
func Initialize() error {
	// Initialize database connection
	if err := initializeDatabase(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Run migrations
	if err := migrations.RunMigrations(DB); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Application initialized successfully")
	return nil
}

// initializeDatabase sets up the database connection
func initializeDatabase() error {
	// Get database connection string from environment variable
	// You can modify this to use your preferred configuration method
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		// Default connection string for local development
		dbURL = "postgres://postgres:password@localhost:5432/simplepay?sslmode=disable"
	}

	var err error
	DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection established")
	return nil
}

// Cleanup performs cleanup operations
func Cleanup() {
	if DB != nil {
		DB.Close()
		log.Println("Database connection closed")
	}
}
