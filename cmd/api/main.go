package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/garvit4540/simplepay/internal/boot"
)

func main() {
	// Initialize the application
	if err := boot.Initialize(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer boot.Cleanup()

	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Application started. Press Ctrl+C to stop.")

	// Wait for shutdown signal
	<-sigChan
	log.Println("Shutting down gracefully...")
}
