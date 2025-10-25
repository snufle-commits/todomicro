package database

import (
	"authservice/internal/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")     // "auth-db"
	dbUser := os.Getenv("DB_USER")     // "authuser"
	dbPass := os.Getenv("DB_PASSWORD") // "authpassword"
	dbName := os.Getenv("DB_NAME")     // "authdbname"
	dbPort := os.Getenv("DB_PORT")     // "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("err to conn tb", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("err to migrate tb", err)
	}

	fmt.Println("Connected to database")
	return db, nil
}
