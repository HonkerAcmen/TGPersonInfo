package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// user, _ := c.Get("user")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "响应成功",
			"data": "",
		})
	}
}
