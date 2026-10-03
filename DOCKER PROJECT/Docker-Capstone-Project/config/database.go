package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = "host=localhost user=registryuser password=registrypass dbname=registrydb port=5432 sslmode=disable"
	}

	var db *gorm.DB
	var err error

	for i := 1; i <= 10; i++ {

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			fmt.Println("Database connected successfully")
			DB = db
			return
		}

		fmt.Printf("Database connection attempt %d failed. Retrying...\n", i)

		time.Sleep(5 * time.Second)
	}

	log.Fatal("Failed to connect database:", err)
}
