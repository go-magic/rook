package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
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
	Addr        string        `yaml:"addr"`
	Passwd      string        `yaml:"passwd"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeout time.Duration `yaml:"idleTimeout"`
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

func (c Config) GetMaxIdle() int {
	return c.Redis.MaxIdle
}

func (c Config) GetMaxActive() int {
	return c.Redis.MaxActive
}

func (c Config) GetIdleTimeout() time.Duration {
	return c.Redis.IdleTimeout
}
