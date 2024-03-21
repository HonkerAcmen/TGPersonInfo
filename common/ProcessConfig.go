package common

import (
	"TGPersonInfo/config"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var _config *config.Config

const SettingFilePath string = "config/setting.yaml"

func InitConfig() {
	c := &config.Config{}
	yamlSic, err := ioutil.ReadFile(SettingFilePath)
	if err != nil {
		panic("获取setting.yaml文件失败 : " + err.Error())
	}
	err = yaml.Unmarshal(yamlSic, c)
	if err != nil {
		panic("解析配置文件失败")
	}
	log.Println("配置文件解析成功")

	_config = c
}

func GetConfig() *config.Config {
	return _config
}
