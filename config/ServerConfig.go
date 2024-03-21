package config

import "fmt"

type ServerConf struct {
	Host  string `yaml:"host"`
	Port  int    `yaml:"port"`
	Debug bool   `yaml:"debug"`
}

func (s ServerConf) GetAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
