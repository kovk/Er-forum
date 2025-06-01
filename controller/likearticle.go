package controller

import (
	"github.com/gin-gonic/gin"
	"newgo/global"
)

func LikeArticle(c *gin.Context) {
	id := c.Param("id")
	likekey := "article:" + id + ":likes"
	if err := global.RedisClient.Incr(likekey).Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "liked"})
}
