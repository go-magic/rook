package auth

import (
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	registerCenter *register
	once           sync.Once
)

type HandlerFunc func(*Auth, *gin.Context)

func GetRegisterInstance() *register {
	once.Do(func() {
		registerCenter = &register{}
		registerCenter.center = make(map[LoginType]HandlerFunc, 3)
	})
	return registerCenter
}

type register struct {
	center map[LoginType]HandlerFunc
}

func (r register) LoginRegister(loginType LoginType, handlerFunc HandlerFunc) {
	r.center[loginType] = handlerFunc
}

func (r register) GetRegister(loginType LoginType) HandlerFunc {
	return r.center[loginType]
}
