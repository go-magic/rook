package auth

import "github.com/gin-gonic/gin"

type Register struct {
	center map[LoginType]gin.HandlerFunc
}

func (r Register) LoginRegister(loginType LoginType, handlerFunc gin.HandlerFunc) {
	r.center[loginType] = handlerFunc
}

func (r Register) GetRegister(loginType LoginType) gin.HandlerFunc {
	return r.center[loginType]
}
