package controller

import (
	"github.com/gin-gonic/gin"
	"newgo/global"
	"newgo/models"
	"time"
)

func CreateExchangeRate(c *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := c.ShouldBindJSON(&exchangeRate); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	exchangeRate.Date = time.Now()
	if err := global.DB.AutoMigrate(&exchangeRate); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&exchangeRate).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, exchangeRate)
}
