package config

import (
	"github.com/go-magic/rook/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var conf *Config

type Config struct {
	config.Common `yaml:"common"`
	Server        `yaml:"server"`
	DB            `yaml:"db"`
}

type Server struct {
	MidServer MidServer `yaml:"midServer"`
}

type DB struct {
	Mysql `yaml:"mysql"`
	Redis `yaml:"redis"`
}

type MidServer struct {
	LogPath string `yaml:"logPath"`
	Debug   bool   `yaml:"debug"`
	Port    string `yaml:"port"`
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

func (c Config) GetLogPath() string {
	return c.MidServer.LogPath
}

func (c Config) GetDebug() bool {
	return c.MidServer.Debug
}

func (c Config) GetPort() string {
	return c.MidServer.Port
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
