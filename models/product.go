package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"image_url"`
}
