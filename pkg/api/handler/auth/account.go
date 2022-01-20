package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/database/mysql/user"
	"github.com/go-magic/rook/pkg/api/handler"
	"net/http"
)

func Account(ctx *gin.Context) {
	auth := &Auth{}
	if err := ctx.BindJSON(auth); err != nil {
		ctx.JSON(http.StatusBadRequest, handler.NewResponse("未知错误", Response{}))
		return
	}
	//check user valid
	sysUser, err := user.GetUserByUsername(auth.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handler.NewResponse("用户不存在", Response{}))
		return
	}
	if sysUser.UserName != auth.Username ||
		sysUser.Passwd != user.Encryption(auth.Passwd) {
		ctx.JSON(http.StatusOK, handler.NewResponse("账号或密码错误", Response{}))
		return
	}
	token, err := claims.CreateToken(sysUser.ID,
		sysUser.UserName, ctx.ClientIP())
	if err != nil {
		ctx.JSON(http.StatusOK, handler.NewResponse("login error", Response{}))
		return
	}
	if token == "" {
		ctx.JSON(http.StatusOK, handler.NewResponse("login error", Response{}))
		return
	}
	if err = setTokenToRedis(sysUser.ID, token); err != nil {
		ctx.JSON(http.StatusOK, handler.NewResponse("login error", Response{}))
		return
	}
	ctx.JSON(http.StatusOK, handler.NewResponse("login error", Response{Token: token}))
}
