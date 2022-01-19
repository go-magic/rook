package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/database/mysql/user"
	"net/http"
)

type Auth struct {
	UserId      uint64 `json:"user_id"`
	Username    string `json:"username"`
	Passwd      string `json:"passwd"`
	PhoneNumber string `json:"phone_number"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	auth := &Auth{}
	if err := ctx.BindJSON(auth); err != nil {
		ctx.JSON(http.StatusBadRequest, NewResponse("", AuthResponse{}))
		return
	}
	//check user valid
	sysUser, err := user.GetUserByUserId(auth.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewResponse("user invalid", AuthResponse{}))
		return
	}
	if sysUser.Passwd != auth.Passwd {
		ctx.JSON(http.StatusOK, NewResponse("user not exist", AuthResponse{}))
		return
	}
	token, err := claims.CreateToken(sysUser.ID, sysUser.UserName)
	if err != nil {
		ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{}))
		return
	}
	ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{Token: token}))
}
