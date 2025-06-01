package controller

import (
	"github.com/gin-gonic/gin"
	"newgo/global"
	"newgo/models"
)

func GetExchangeRate(c *gin.Context) {
	var exchangeRates []models.ExchangeRate
	if err := global.DB.Find(&exchangeRates).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, exchangeRates)
}
