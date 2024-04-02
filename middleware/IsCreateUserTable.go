package middleware

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"

	"github.com/gin-gonic/gin"
)

func IsCreateUserTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := model.GetDB()
		logger := logging.GetLog()
		if !db.Migrator().HasTable("userinfos") {
			db.AutoMigrate(&model.UserInfo{})
			// logger.Error("自动创表失败", err)
			logger.Info("检查表成功")
		}
	}
}
