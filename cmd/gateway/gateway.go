package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/config"
	"github.com/go-magic/rook/pkg/api/database/mysql"
	"github.com/go-magic/rook/pkg/api/database/mysql/user"
	"github.com/go-magic/rook/pkg/api/database/redis"
	"github.com/go-magic/rook/pkg/api/handler/auth"
	"github.com/go-magic/rook/pkg/api/middleware"
	"github.com/go-magic/rook/pkg/api/router"
)

func initRouter() error {
	r := router.NewRouter()
	login(r)
	authorization(r)
	return r.Run()
}

func login(r *gin.Engine) {
	auth.GetRegisterInstance().LoginRegister(auth.USER_PASSWD, auth.Account)
	auth.GetRegisterInstance().LoginRegister(auth.PHONE_NUMBER, auth.Phone)
	auth.GetRegisterInstance().LoginRegister(auth.EMAIL_PASSWD, auth.Email)
	r.Handle("POST", "/api/login", auth.Login)
}

func authorization(r *gin.Engine) {
	r.Handle("POST", "/api/logout", auth.Logout).Use(middleware.Authorization)
}

func initConfig() error {
	return config.NewConfig("./etc/config/config.yml")
}

func initMysql() error {
	if err := mysql.InitMysql(config.GetConfig().GetMysqlAddr()); err != nil {
		return err
	}
	return mysql.GetMysqlInstance().InitTables(
		user.User{},
	)
}

func initRedis() error {
	return redis.InitRedisPool(
		config.GetConfig().GetRedisAddr(),
		config.GetConfig().GetMaxIdle(),
		config.GetConfig().GetMaxActive(),
		config.GetConfig().GetIdleTimeout(),
		config.GetConfig().GetRedisPasswd())
}

func startServer() error {
	if err := initConfig(); err != nil {
		return err
	}
	if err := initMysql(); err != nil {
		return err
	}
	if err := initRedis(); err != nil {
		return err
	}
	if err := initRouter(); err != nil {
		return err
	}
	return errors.New("程序退出")
}

func main() {
	if err := startServer(); err != nil {
		panic(err)
	}
}
