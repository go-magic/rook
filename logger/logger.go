package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Args struct {
	Error  error
	Fields []zap.Field
}

func InitLogger(logPath string, debug bool) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	}
	level := Level(debug)
	// 设置日志级别
	atom := zap.NewAtomicLevelAt(level)

	config := zap.Config{
		Level:            atom,                                             // 日志级别
		Development:      false,                                            // 开发模式，堆栈跟踪
		Encoding:         "console",                                        // 输出格式 console 或 json
		EncoderConfig:    encoderConfig,                                    // 编码器配置
		InitialFields:    map[string]interface{}{"serviceName": "monitor"}, // 初始化字段，如：添加一个服务器名称
		OutputPaths:      []string{"stdout", logPath},                      // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
		ErrorOutputPaths: []string{"stderr"},
	}
	var err error
	// 构建日志
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("logger 初始化失败: %v", err))
	}
	logger.Info("logger 初始化成功")
}

/*
Level 设置日志等级
*/
func Level(debug bool) zapcore.Level {
	if debug {
		return zapcore.DebugLevel
	}
	return zapcore.InfoLevel
}

var logger = &zap.Logger{}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

//Error 固定格式
func Error(msg string, args, err zap.Field) {
	logger.Warn(msg, args, err)
}

func ErrorF(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
