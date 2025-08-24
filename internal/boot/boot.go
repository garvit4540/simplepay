package boot

import (
	"fmt"
	"github.com/garvit4540/simplepay/internal/database"
	_ "github.com/lib/pq"
	"log"
)

// Initialize initializes the application
func Initialize() error {
	if err := database.InitializeDatabase(); err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	if err := database.RunMigrations(); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Application initialized successfully")
	return nil
}

// Cleanup performs cleanup operations
func Cleanup() {

}
