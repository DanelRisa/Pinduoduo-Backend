package database

import (
	"fmt"
	"log"
	"pinduoduo-back/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=251046dd dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect :", err)
	}
	fmt.Println("Database connected ")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Product{}, &models.GroupBuy{},
		&models.Order{},)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
	fmt.Println("Database migrated ")
}
