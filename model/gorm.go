package model

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/common"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitDB() {
	var config = common.GetConfig()
	var Logger = logging.GetLog()
	if config.Mysql.Host == "" {
		Logger.Fatalln("数据库Host为空")
	}
	dsn := config.Mysql.DSN()
	var err error
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		Logger.Fatalln("数据库连接失败")
	}
	sql_db, _ := _db.DB()
	sql_db.SetMaxIdleConns(20)
	sql_db.SetMaxOpenConns(100)
	sql_db.SetConnMaxLifetime(time.Hour * 4)
	Logger.Info("数据库连接成功")

}

func GetDB() *gorm.DB {
	return _db
}
