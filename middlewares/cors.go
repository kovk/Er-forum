package middlewares

import (
	"github.com/gin-gonic/gin"
	"newgo/config"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取请求方法

		//添加跨域响应头
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", config.AppConfig.App.Host+config.AppConfig.App.Port)
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "X-Token,Sessionid, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法

		//处理请求
		c.Next()
	}
}
