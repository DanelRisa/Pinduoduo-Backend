package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB        // Для основной базы данных
var TestDB *gorm.DB    // Для тестовой базы данных

// Connect — подключение к основной базе данных
func Connect() error {
	dsn := "host=db user=postgres password=251046dd dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return err
	}
	fmt.Println("Connected to the main database.")
	return nil
}

// ConnectTestDB — подключение к тестовой базе данных
func ConnectTestDB() error {
	dsn := "host=localhost user=postgres password=251046dd dbname=testdb port=5432 sslmode=disable"
	var err error
	TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the test database:", err)
		return err
	}
	fmt.Println("Connected to the test database.")
	return nil
}