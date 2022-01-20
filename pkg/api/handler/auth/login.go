package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/database/redis"
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

type Response struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {

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
