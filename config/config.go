package config

type Config struct {
	Common `yaml:"common"`
}

type Common struct {
	MysqlAddr string `yaml:"mysqlAddr"`
}
