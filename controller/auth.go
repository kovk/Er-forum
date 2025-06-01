package controller

import (
	"github.com/gin-gonic/gin"
	"newgo/global"
	"newgo/models"
	"newgo/utils"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	hashpwd, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashpwd
	signedToken, err := utils.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.AutoMigrate(&user); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if err := global.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": signedToken})

}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.DB.Where("Username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	if !utils.CheckLogin(user.Password, input.Password) {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	} else {
		signedToken, err := utils.GenerateJWT(user.Username)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"token": signedToken})
	}
}
