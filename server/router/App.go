package router

import (
	"github.com/gin-gonic/gin"
	"shopping_mall/service"
)

func App() *gin.Engine {
	r := gin.Default()

	// 使用默认配置的CORS中间件
	// 默认允许所有来源（AllowAllOrigins: true）
	//	r.Use(cors.Default())

	r.Use(CORSMiddleware())

	r.POST("/login/psw", service.LoginPsw)

	return r
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// http://localhost:5173  是前端vue项目的node启动的地址和端口
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
