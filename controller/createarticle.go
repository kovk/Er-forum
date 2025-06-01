package controller

import (
	"github.com/gin-gonic/gin"
	"newgo/global"
	"newgo/models"
	"time"
)

func CreateArticle(c *gin.Context) {
	var article models.Article
	article.Date = time.Now()
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	if err := global.DB.AutoMigrate(&article); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	if err := global.DB.Create(&article).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Article created successfully"})
	global.RedisClient.Del(cachekey)
}
