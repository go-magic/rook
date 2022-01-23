package claims

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-magic/rook/pkg/api/database/redis"
	"time"
)

var Secret = []byte("as1%^&#!*$%secret&^") // 用来加密解密

type Claims struct {
	UserID   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateToken(userID uint64, username string, ip string) (string, error) {
	var claims = Claims{
		userID,
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(redis.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "***",                                            // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", fmt.Errorf("生成token失败:%v", err)
	}
	return signedToken, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
		func(token *jwt.Token) (i interface{}, err error) { // 解析token
			return Secret, nil
		})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func Valid(token string) bool {
	_, err := ParseToken(token)
	if err != nil {
		return false
	}
	return true
}
