package middlewares

import (
	"github.com/gin-gonic/gin"
	"newgo/utils"
)

func Authmiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		username, err := utils.ParseJWT(token)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
		}
		c.Set("username", username)
		c.Next()
	}
}
