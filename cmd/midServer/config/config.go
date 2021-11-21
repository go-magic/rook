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
}

type Server struct {
	MidServer MidServer `yaml:"midServer"`
}

type MidServer struct {
	LogPath string `yaml:"logPath"`
	Debug   bool   `yaml:"debug"`
	Port    string `yaml:"port"`
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
	return c.MysqlAddr
}
