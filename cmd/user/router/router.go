package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/cmd/midServer/router/handler"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	initUserHandle(router)
	return router
}

func initUserHandle(router *gin.Engine) {
	router.Handle("POST", "/api/register", handler.AddTask)
}
