package common

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/model"
)

func AutoCreateTable(userinfo model.UserInfo) {
	db := model.GetDB()
	logger := logging.GetLog()
	if err := db.AutoMigrate(&userinfo).Error(); err != "" {
		logger.Error("自动创表失败", err)
	}
	logger.Info("创表成功")
}
