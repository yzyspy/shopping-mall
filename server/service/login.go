package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPsw(c *gin.Context) {
	request := new(LoginPswRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
}
