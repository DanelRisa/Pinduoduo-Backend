package controllers

import (
	"net/http"
	// "strconv"

	"github.com/gin-gonic/gin"
	"pinduoduo-back/models"
	"pinduoduo-back/database"
)

func CreateGroupBuy(c *gin.Context) {
	var groupbuy models.GroupBuy

	if err := c.ShouldBindJSON(&groupbuy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupbuy.Status = "active"
	groupbuy.Participants = 0

	if err := database.DB.Create(&groupbuy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}

	c.JSON(http.StatusCreated, groupbuy)
}

func GetGroupBuys(c *gin.Context) {
	var groupbuys []models.GroupBuy
	if err := database.DB.Preload("Product").Find(&groupbuys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении"})
		return
	}
	c.JSON(http.StatusOK, groupbuys)
}

func GetGroupBuy(c *gin.Context) {
	id := c.Param("id")
	var groupbuy models.GroupBuy

	if err := database.DB.Preload("Product").First(&groupbuy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Групповая покупка не найдена"})
		return
	}

	c.JSON(http.StatusOK, groupbuy)
}

func JoinGroupBuy(c *gin.Context) {
	id := c.Param("id")
	var groupbuy models.GroupBuy

	if err := database.DB.First(&groupbuy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Групповая покупка не найдена"})
		return
	}

	if groupbuy.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Покупка неактивна"})
		return
	}

	groupbuy.Participants += 1

	// Если участников достаточно закрываем
	if groupbuy.Participants >= groupbuy.MinParticipants {
		groupbuy.Status = "closed"
	}

	if err := database.DB.Save(&groupbuy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, groupbuy)
}
