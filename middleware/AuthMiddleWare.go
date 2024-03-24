package middleware

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/common"
	"TGPersonInfo/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthmiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := logging.GetLog()
		db := model.GetDB()
		text := "TGP"
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 510,
				"msg":  "token为空",
			})
			c.Abort()
			return
		}
		index := strings.Index(tokenString, text+":")
		tokenString = tokenString[index+len(text)+1:]

		jwtToken, claims, err := common.ParseToken(tokenString)
		if err != nil || !jwtToken.Valid {
			logger.Error("权限认证组件中Token解析错误或等待过期", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 511,
				"msg":  "权限不足",
			})
			c.Abort()
			return
		}

		userId := claims.UserID
		var user model.UserInfo
		if err := db.First(&user, userId).Error; err != nil {
			logger.Error("认证组件中为查到用户", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 512,
				"msg":  "认证组件未找到用户",
			})
			c.Abort()
			return
		}

		logger.Info("权限通过")
		c.Set("user", user)
		c.Next()
	}
}
