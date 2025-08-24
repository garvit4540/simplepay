package migrations

import (
	"database/sql"
	"fmt"
	"sort"
)

// Migration represents a database migration
type Migration struct {
	Version int
	Up      func(*sql.Tx) error
	Down    func(*sql.Tx) error
}

// Migrations slice holds all migrations
var Migrations []*Migration

// RunMigrations runs all pending migrations
func RunMigrations(db *sql.DB) error {
	// Create migrations table if it doesn't exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version INT PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create schema_migrations table: %w", err)
	}

	// Get applied migrations
	appliedVersions := make(map[int]bool)
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return fmt.Errorf("failed to query applied migrations: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var version int
		if err := rows.Scan(&version); err != nil {
			return fmt.Errorf("failed to scan migration version: %w", err)
		}
		appliedVersions[version] = true
	}

	// Sort migrations by version
	sort.Slice(Migrations, func(i, j int) bool {
		return Migrations[i].Version < Migrations[j].Version
	})

	// Run pending migrations
	for _, migration := range Migrations {
		if !appliedVersions[migration.Version] {
			fmt.Printf("Running migration version %d\n", migration.Version)
			
			tx, err := db.Begin()
			if err != nil {
				return fmt.Errorf("failed to begin transaction for migration %d: %w", migration.Version, err)
			}

			if err := migration.Up(tx); err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to run migration %d: %w", migration.Version, err)
			}

			// Record migration as applied
			_, err = tx.Exec("INSERT INTO schema_migrations (version) VALUES (?)", migration.Version)
			if err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to record migration %d: %w", migration.Version, err)
			}

			if err := tx.Commit(); err != nil {
				return fmt.Errorf("failed to commit migration %d: %w", migration.Version, err)
			}

			fmt.Printf("Successfully applied migration version %d\n", migration.Version)
		}
	}

	return nil
}
