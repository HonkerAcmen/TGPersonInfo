package config

type Config struct {
	Server ServerConf `yaml:"server"`
	Mysql  MysqlConf  `yaml:"mysql"`
	Log    LogConf    `yaml:"log"`
}
