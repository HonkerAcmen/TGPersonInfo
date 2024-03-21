package config

type LogConf struct {
	Projectname string `yaml:"projectname"`
	Level       string `yaml:"level"`
	Showline    bool   `yaml:"showline"`
	IsLog       bool   `yaml:"isLog"`
}
