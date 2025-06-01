package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"newgo/global"
)

func GetArticleLikes(c *gin.Context) {
	articleID := c.Param("id")
	likekey := "article:" + articleID + ":likes"
	likes, err := global.RedisClient.Get(likekey).Result()
	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"likes": likes})

}
