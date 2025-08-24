package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

func RunMigrations() error {
	db, err := DatabaseClient.DB()
	if err != nil {
		return fmt.Errorf("failed to covert gorm client to sql client")
	}

	if err := goose.Up(db, "internal/database/migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("✅ Database migrations applied successfully")
	return nil
}
