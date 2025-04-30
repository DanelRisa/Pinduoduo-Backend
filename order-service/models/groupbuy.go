// models/groupbuy.go
package models

import "gorm.io/gorm"

type GroupBuy struct {
	gorm.Model
	ProductID       uint    `json:"product_id"`
	Product         Product `gorm:"foreignKey:ProductID"` //
	Discount        float64 `json:"discount"`
	Participants    int     `json:"participants"`
	MinParticipants int     `json:"min_participants"`
	Status          string  `json:"status"`
}
