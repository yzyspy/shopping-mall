package router

import (
	"github.com/gin-gonic/gin"
	"shopping_mall/service"
)

func App() *gin.Engine {
	r := gin.Default()

	r.POST("/login/psw", service.LoginPsw)

	return r
}
