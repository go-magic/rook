package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/database/mysql/user"
	"github.com/go-magic/rook/pkg/api/database/redis"
	"net/http"
	"time"
)

var MD5Secret = "secret" // 用来加密解密

type LoginType int

const (
	USER_PASSWD LoginType = iota
	PHONE_NUMBER
	EMAIL_PASSWD
)

type Auth struct {
	UserId      uint64 `json:"user_id"`
	LoginType   int    `json:"login_type"`
	Username    string `json:"username"`
	Passwd      string `json:"passwd"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
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
	sysUser, err := user.GetUserByUsername(auth.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, NewResponse("用户不存在", AuthResponse{}))
		return
	}
	if sysUser.UserName != auth.Username ||
		sysUser.Passwd != user.Encryption(auth.Passwd) {
		ctx.JSON(http.StatusOK, NewResponse("账号或密码错误", AuthResponse{}))
		return
	}
	token, err := claims.CreateToken(sysUser.ID,
		sysUser.UserName, ctx.ClientIP())
	if err != nil {
		ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{}))
		return
	}
	if token == "" {
		ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{}))
		return
	}
	if err = setTokenToRedis(sysUser.ID, token); err != nil {
		ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{}))
		return
	}
	ctx.JSON(http.StatusOK, NewResponse("login error", AuthResponse{Token: token}))
}

func getTokenByRedis(userId uint64) (string, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := redis.GetPool().GetContext(ctx)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = conn.Close()
	}()
	return redis.GetString(conn, userId)
}

func setTokenToRedis(userId uint64, token string) error {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := redis.GetPool().GetContext(ctx)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Do("SETEX", userId, redis.TokenExpireDuration, token)
	return err
}
