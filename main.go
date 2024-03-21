package main

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/common"
	"TGPersonInfo/model"
)

func main() {
	common.InitConfig()
	logging.InitLogger()
	model.InitDB()
}
