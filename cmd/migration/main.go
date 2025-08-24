package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/garvit4540/simplepay/internal/boot"
	"github.com/garvit4540/simplepay/internal/database/migrations"
)

func main() {
	// Parse command line flags
	var (
		up   = flag.Bool("up", false, "Run migrations up")
		down = flag.Bool("down", false, "Run migrations down")
	)
	flag.Parse()

	// Initialize database connection
	if err := boot.Initialize(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer boot.Cleanup()

	if *up {
		fmt.Println("Running migrations up...")
		if err := migrations.RunMigrations(boot.DB); err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Migrations completed successfully")
	} else if *down {
		fmt.Println("Running migrations down...")
		// Note: Down migrations are not implemented in this simple version
		// You would need to implement a more sophisticated migration system
		// that tracks which migrations to rollback
		fmt.Println("Down migrations not implemented yet")
	} else {
		fmt.Println("Usage: go run cmd/migration/main.go -up")
		fmt.Println("  -up: Run migrations up")
		fmt.Println("  -down: Run migrations down (not implemented)")
		os.Exit(1)
	}
}
