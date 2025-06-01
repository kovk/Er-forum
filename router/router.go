package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"newgo/controller"
	"newgo/middlewares"
	"time"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", controller.Login)
		auth.POST("/register", controller.Register)
	}
	api := router.Group("/api")
	api.GET("/exchangeRates", controller.GetExchangeRate)
	api.Use(middlewares.Authmiddleware())
	{
		api.POST("/exchangeRates", controller.CreateExchangeRate)
		api.POST("/articles", controller.CreateArticle)
		api.GET("/articles", controller.GetAllArticle)
		api.GET("/articles/:id", controller.GetArticleById)
		api.POST("/articles/:id/like", controller.LikeArticle)
		api.GET("/articles/:id/like", controller.GetArticleLikes)
	}

	return router
}
