package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping_mall/models"
)

func LoginPsw(c *gin.Context) {
	request := new(LoginPswRequest)
	fmt.Println("LoginPsw request:", request)
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
	}
	u, err := models.GetUserByUserNameAndPsw(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  err.Error(),
		})
		return
	}
	if u == nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "user not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "login success",
	})
}
