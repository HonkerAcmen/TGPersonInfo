package main

import (
	"TGPersonInfo/common"
	"TGPersonInfo/config"
	"fmt"
)

func main() {
	common.InitConfig()
	c := common.GetConfig()
	fmt.Println(config.ServerConf.GetAddress(c.Server))
}
