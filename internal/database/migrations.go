package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func RunMigrations() error {
	db, err := DatabaseClient.DB()
	if err != nil {
		return fmt.Errorf("failed to convert gorm client to sql client")
	}

	// Get migration files
	migrationDir := "internal/database/migrations"
	files, err := os.ReadDir(migrationDir)
	if err != nil {
		return fmt.Errorf("failed to read migration directory: %w", err)
	}

	// Filter and sort SQL files
	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	// Execute each migration file
	for _, fileName := range sqlFiles {
		filePath := filepath.Join(migrationDir, fileName)
		content, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", fileName, err)
		}

		// Extract UP migration (everything between -- +goose Up and -- +goose Down)
		contentStr := string(content)
		upStart := strings.Index(contentStr, "-- +goose Up")
		downStart := strings.Index(contentStr, "-- +goose Down")

		if upStart == -1 || downStart == -1 {
			log.Printf("Warning: Skipping %s - invalid migration format", fileName)
			continue
		}

		upMigration := contentStr[upStart:downStart]
		// Remove the -- +goose Up line and extract SQL statements
		lines := strings.Split(upMigration, "\n")
		var sqlStatements []string
		var currentStatement strings.Builder
		inStatement := false

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "-- +goose Up" {
				continue
			}
			if line == "-- +goose StatementBegin" {
				inStatement = true
				continue
			}
			if line == "-- +goose StatementEnd" {
				if currentStatement.Len() > 0 {
					sqlStatements = append(sqlStatements, currentStatement.String())
					currentStatement.Reset()
				}
				inStatement = false
				continue
			}
			if inStatement && line != "" {
				currentStatement.WriteString(line)
				currentStatement.WriteString("\n")
			}
		}

		// Execute each SQL statement
		for _, statement := range sqlStatements {
			statement = strings.TrimSpace(statement)
			if statement != "" {
				log.Printf("Executing migration: %s", fileName)
				_, err := db.Exec(statement)
				if err != nil {
					// Check if it's a "table already exists" or "duplicate key" error and skip it
					if strings.Contains(err.Error(), "already exists") || strings.Contains(err.Error(), "Duplicate key") {
						log.Printf("Object already exists, skipping: %s", fileName)
						continue
					}
					return fmt.Errorf("failed to execute migration %s: %w", fileName, err)
				}
			}
		}
	}

	log.Println("✅ Database migrations applied successfully")
	return nil
}
