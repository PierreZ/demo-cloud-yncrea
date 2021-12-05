package models

import (
	"fmt"
	"strconv"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabaseWithEnvVars() error {
	dbName := os.Getenv("POSTGRESQL_ADDON_DB")
	host := os.Getenv("POSTGRESQL_ADDON_HOST")
	password := os.Getenv("POSTGRESQL_ADDON_PASSWORD")
	user := os.Getenv("POSTGRESQL_ADDON_USER")
	sslMode := os.Getenv("POSTGRESQL_ADDON_SSLMODE")
	port, err := strconv.Atoi(os.Getenv("POSTGRESQL_ADDON_PORT"))
	if err != nil {
		return fmt.Errorf("could not parse pg port: %w", err)
	}
	return connectDatabase(host, user, password, dbName, port, sslMode)
}

func connectDatabase(host string, user string, password string, dbName string, port int, sslMode string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", host, user, password, dbName, port, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("could not connect to DB: %w", err)
	}
	DB = db

	err = db.AutoMigrate(&Book{})
	if err != nil {
		return fmt.Errorf("could not migrate db book: %w", err)
	}

	DB.Create(&Book{
		ID:     1,
		Title:  "The Hitchhiker's Guide to the Galaxy",
		Author: "Douglas Adams",
	})

	if err != nil {
		return fmt.Errorf("could not create initial book: %w", err)
	}

	return nil
}
