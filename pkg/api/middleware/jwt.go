package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/database/redis"
	"github.com/go-magic/rook/pkg/api/handler"
	"net/http"
	"time"
)

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, handler.NewResponse("请登录", nil))
		return
	}
	if !claims.Valid(token) {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, handler.NewResponse("请登录", nil))
		return
	}
	claimsToken, err := claims.ParseToken(token)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, handler.NewResponse("请登录", nil))
		return
	}
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := redis.GetPool().GetContext(timeout)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, handler.NewResponse("服务器内部异常", nil))
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	redisToken, err := redis.GetString(conn, claimsToken.UserID)
	if err != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, handler.NewResponse("服务器内部异常", nil))
		return
	}
	if token != redisToken {
		ctx.Abort()
		ctx.JSON(http.StatusBadRequest, handler.NewResponse("非法用户", nil))
		return
	}
	ctx.Next()
}
