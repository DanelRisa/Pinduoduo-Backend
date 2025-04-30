package controllers

import (
	"net/http"
	"pinduoduo-back/database"
	"pinduoduo-back/order-service/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создании заказа"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := database.DB.Preload("GroupBuy").Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получении заказов"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	var order models.Order
	orderID := c.Param("id")

	if err := database.DB.Preload("GroupBuy").Preload("Product").First(&order, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving order"})
		return
	}

	c.JSON(http.StatusOK, order)
}


func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправ ID"})
		return
	}

	if err := database.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удалении заказа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заказ удалён"})
}
