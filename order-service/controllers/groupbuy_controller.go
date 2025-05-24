package controllers

import (
	"net/http"

	"pinduoduo-back/database"
	"pinduoduo-back/order-service/models"

	"github.com/gin-gonic/gin"
)

func CreateGroupBuy(c *gin.Context) {
	var groupbuy models.GroupBuy

	if err := c.ShouldBindJSON(&groupbuy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var product models.Product
	if err := database.DB.First(&product, groupbuy.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
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


//

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
		c.JSON(http.StatusNotFound, gin.H{"error": "ошибка айди"})
		return
	}

	c.JSON(http.StatusOK, groupbuy)
}

func JoinGroupBuy(c *gin.Context) {
	id := c.Param("id")
	var groupbuy models.GroupBuy

	if err := database.DB.First(&groupbuy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "не найдена"})
		return
	}

	if groupbuy.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка айди"})
		return
	}

	groupbuy.Participants += 1

	// закр если мдостног
	if groupbuy.Participants >= groupbuy.MinParticipants {
		groupbuy.Status = "closed"
	}

	if err := database.DB.Save(&groupbuy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновлении"})
		return
	}

	c.JSON(http.StatusOK, groupbuy)
}
func DeleteGroupBuy(c *gin.Context) {
	id := c.Param("id")
	var groupbuy models.GroupBuy
	if err := database.DB.First(&groupbuy, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "groupbuy not found"})
		return
	}

	if err := database.DB.Delete(&groupbuy).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "groupbuy deleted"})
}
