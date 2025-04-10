package controllers

import (
	"net/http"
	"pinduoduo-back/database"
	"pinduoduo-back/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	if err := database.DB.Preload("GroupBuy").Preload("Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении заказов"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := database.DB.Preload("GroupBuy").Preload("Product").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	orderID, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := database.DB.Delete(&models.Order{}, orderID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении заказа"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заказ удалён"})
}
