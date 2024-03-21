package config

import "fmt"

type MysqlConf struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Databasename string `yaml:"databasename"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Charset      string `yaml:"charset"`
}

func (sql MysqlConf) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		sql.User,
		sql.Password,
		sql.Host,
		sql.Port,
		sql.Databasename,
		sql.Charset,
	)
}
