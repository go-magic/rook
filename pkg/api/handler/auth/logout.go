package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/database/redis"
	"github.com/go-magic/rook/pkg/api/handler"
	"net/http"
	"time"
)

func Logout(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	claimsToken, err := claims.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, handler.NewResponse("请登录", nil))
		return
	}
	timeout, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := redis.GetPool().GetContext(timeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, handler.NewResponse("服务器内部异常", nil))
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	_, err = conn.Do("del", claimsToken.UserID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, handler.NewResponse("请登录", nil))
		return
	}
	ctx.JSON(http.StatusOK, handler.NewResponse("退出登录", nil))
}
