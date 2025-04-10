package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ProductID  uint     `json:"product_id"`
	Product    Product  `gorm:"foreignKey:ProductID"`
	GroupBuyID uint     `json:"groupbuy_id"`
	GroupBuy   GroupBuy `gorm:"foreignKey:GroupBuyID"`
	Quantity   int      `json:"quantity"`
	TotalPrice float64  `json:"total_price"`
}
