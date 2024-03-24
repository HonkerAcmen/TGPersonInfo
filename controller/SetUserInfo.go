package controller

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"
	"TGPersonInfo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userinfo model.UserInfo
		db := model.GetDB()
		logger := logging.GetLog()
		err := c.Bind(&userinfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 510,
				"msg":  "绑定失败",
			})
			return
		}

		newUser := model.UserInfo{
			UserID:    utils.RandomInt(10),
			UserName:  userinfo.UserName,
			UserSign:  userinfo.UserSign,
			Following: userinfo.Following,
			Followers: userinfo.Followers,
		}
		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 510,
				"msg":  "设置信息失败",
			})
			logger.Error(err)
			return
		}

	}
}
