package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-magic/rook/pkg/api/claims"
	"github.com/go-magic/rook/pkg/api/handler"
	"net/http"
)

func Authorization(ctx *gin.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, handler.Response{})
		return
	}
	if !claims.Valid(token) {
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, handler.Response{})
		return
	}
	ctx.Next()
}
