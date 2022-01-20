package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var conf *Config

type Config struct {
	Common `yaml:"common"`
	DB     `yaml:"db"`
}

type Common struct {
}

type DB struct {
	Mysql `yaml:"mysql"`
	Redis `yaml:"redis"`
}

type Mysql struct {
	Addr string `yaml:"addr"`
}

type Redis struct {
	Addr   string `yaml:"addr"`
	Passwd string `yaml:"passwd"`
}

func NewConfig(path string) error {
	conf = &Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yamlFile, conf)
}

func GetConfig() *Config {
	return conf
}

func (c Config) GetMysqlAddr() string {
	return c.Mysql.Addr
}

func (c Config) GetRedisAddr() string {
	return c.Redis.Addr
}

func (c Config) GetRedisPasswd() string {
	return c.Redis.Passwd
}
