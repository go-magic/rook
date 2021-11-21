package main

import (
	"github.com/go-magic/rook/cmd/user/config"
	"github.com/go-magic/rook/cmd/user/router"
	"github.com/go-magic/rook/logger"
	"github.com/go-magic/rook/pkg/api/database/mysql"
	"github.com/go-magic/rook/pkg/api/database/mysql/user"
	"go.uber.org/zap"
)

const configPath = "./etc/config/config.yml"

func initConfig(path string) error {
	return config.NewConfig(path)
}

func initLog(path string, debug bool) {
	logger.InitLogger(path, debug)
}

func initServer() {
	if err := initConfig(configPath); err != nil {
		panic(err)
	}
	initLog(config.GetConfig().GetLogPath(), config.GetConfig().GetDebug())
	if err := initMysql(config.GetConfig().GetMysqlAddr()); err != nil {
		panic(err)
	}
}

func initMysql(mysqlAddr string) error {
	if err := mysql.InitMysql(mysqlAddr); err != nil {
		return err
	}
	return mysql.GetMysqlInstance().InitTables(user.User{})
}

func startServer() {
	initServer()
	logger.Fatal("gin退出", zap.Any("error", router.NewRouter().Run(":"+config.GetConfig().GetPort())))
}

func main() {
	startServer()
}
