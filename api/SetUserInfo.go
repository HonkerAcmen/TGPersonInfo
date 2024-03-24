package api

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"
	"TGPersonInfo/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// isUserExist: 如果用户存在，则返回true
func isUserExist(db *gorm.DB, username string) bool {
	var user model.UserInfo
	err := db.Where("username=?", username).First(&user).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

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
		if isUserExist(db, userinfo.UserName) {
			logger.Error("用户数据已经存在")
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 500,
				"msg":  "用户数据已经存在",
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
		logger.Info("设置信息成功")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "设置成功",
		})
	}
}
