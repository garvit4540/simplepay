package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DatabaseClient *gorm.DB

func InitializeDatabase() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "simplepay_user:12345678@tcp(127.0.0.1:3306)/simplepay?charset=utf8mb4&parseTime=True&loc=Local"
	}

	// Open connection using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("❌ Failed to connect to MySQL: %v", err)
		return err
	}

	// Assign global DB instance
	DatabaseClient = db
	log.Println("✅ Successfully connected to MySQL database")
	return nil
}
