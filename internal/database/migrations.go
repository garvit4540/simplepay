package database

import (
	"fmt"
	"github.com/pressly/goose"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// RunMigrations applies all migrations in the migrations folder
func RunMigrations() error {
	db, err := DatabaseClient.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm: %w", err)
	}

	if err := goose.SetDialect("mysql"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	migrationsDir := "internal/database/migrations"
	if err := goose.Up(db, migrationsDir); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("✅ Database migrations applied successfully")
	return nil
}
