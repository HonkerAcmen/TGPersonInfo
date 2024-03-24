package api

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userinfo []model.UserInfo
		db := model.GetDB()
		logger := logging.GetLog()

		err := db.Find(&userinfo).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("无任何用户数据")
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 404,
				"msg":  "无任何用户数据",
			})
			return
		}

		logger.Info("查询成功")
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "查询成功",
			"data": userinfo,
		})
	}
}
