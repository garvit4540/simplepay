package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/garvit4540/simplepay/internal/boot"
)

func main() {
	if err := boot.Initialize(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer boot.Cleanup()

	// Setting up signal handling for graceful shut down
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Application started. Press Ctrl+C to stop.")

	<-sigChan
	log.Println("Shutting down gracefully...")
}
