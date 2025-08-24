package main

import (
	"log"
	"os"
	"time"

	"github.com/garvit4540/simplepay/internal/routing"

	"github.com/garvit4540/simplepay/internal/boot"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")

	if err := boot.Initialize(); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	defer clean()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	routing.SetupRouter(r)

	port := os.Getenv("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start on port : %s", port)
	}
	log.Printf("Server running on port %s", port)
}

func clean() {
	err := boot.Cleanup()
	if err != nil {
		log.Fatalf("Failed to gracefully close application: %v", err)
	}
}
