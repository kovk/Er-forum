package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"newgo/global"
	"newgo/models"
)

func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article
	if err := global.DB.Where("id = ?", id).First(&article).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, article)
}
