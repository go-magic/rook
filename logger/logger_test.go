package logger

import (
	"go.uber.org/zap"
	"testing"
)

func TestLog(t *testing.T) {
	InitLogger("./logger.logger", false)
	Warn("请求错误",
		zap.String("url", "http://www.baidu.com"),
		zap.String("error", "请求失败"))
	Debug("123", zap.String("key", "key"))
}
