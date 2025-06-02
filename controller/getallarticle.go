package controller

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"newgo/global"
	"newgo/models"
	"sync"
	"time"
)

var cachekey = "articles"

func GetAllArticle(c *gin.Context) {
	var articles []models.Article
	cachekey = "articles"
	redisResult, err := global.RedisClient.Get(cachekey).Result()
	if err == redis.Nil {
		mu.Lock()
		defer mu.Unlock()
		if err := global.DB.Find(&articles).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(404, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
		}
		articleJSON, err := json.Marshal(articles)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		global.RedisClient.Set(cachekey, articleJSON, 10*time.Minute)
	} else {
		if err := json.Unmarshal([]byte(redisResult), &articles); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(200, articles)
}
