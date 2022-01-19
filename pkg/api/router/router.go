package router

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	return router
}
