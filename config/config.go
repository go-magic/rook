package config

type Config struct {
	Common `yaml:"common"`
}

type Common struct {
}

type Mysql struct {
	Addr string `yaml:"addr"`
}

type Redis struct {
	Addr   string `yaml:"addr"`
	Passwd string `yaml:"passwd"`
}
