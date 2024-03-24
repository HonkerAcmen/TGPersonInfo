package main

import (
	logging "TGPersonInfo/Logging"
	"TGPersonInfo/config"
	"TGPersonInfo/model"
	"TGPersonInfo/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	c := config.GetConfig()
	logging.InitLogger()
	model.InitDB()
	r := gin.Default()
	r = router.CollectRouter(r)

	addr := config.ServerConf.GetAddress(c.Server)

	err := r.Run(addr)
	if err != nil {
		return
	}
}
