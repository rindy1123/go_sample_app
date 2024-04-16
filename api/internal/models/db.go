package models

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		panic("Failed to connect to the database")
	}
	return db
}

// Add the new models to the tables slice
var tables = []interface{}{
	&Todo{},
}

func SetupTestDB() *gorm.DB {
	uuid := uuid.New().String()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		panic("Failed to connect to the database")
	}
	db.Exec(fmt.Sprintf("CREATE DATABASE \"%s\"", uuid))

	dsn = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		uuid,
		os.Getenv("DB_PORT"),
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		panic("Failed to connect to the database")
	}
	migrator := db.Migrator()
	migrator.DropTable(tables...)
	migrator.AutoMigrate(tables...)
	return db
}
