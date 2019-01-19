package logger

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//Field 参数列表
type Field = zapcore.Field

var encoderCfg = zapcore.EncoderConfig{
	TimeKey:        "Time",
	LevelKey:       "Level",
	NameKey:        "Name",
	CallerKey:      "",
	MessageKey:     "Msg",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}
var level = zap.NewAtomicLevel()
var prefix string

// Init 模块初始化
func Init() bool {
	ss := strings.Split(os.Args[0], "/")
	prefix = ss[len(ss)-1]
	timeS := time.Now().Format("2006-01-02")
	cfg := zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{prefix + timeS + ".log", "stdout"},
		ErrorOutputPaths: []string{prefix + timeS + ".log", "stdout"},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		Fatal("logger init error")
		return false
	}
	zap.ReplaceGlobals(logger)
	Info("logger init success")
	return true
}

//LogGin gin的日志封装，不建议用
func LogGin() func(*gin.Context) {
	return func(c *gin.Context) {
		//start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()
		//latency := time.Now().Sub(start)
		mothod := c.Request.Method
		clientIP := c.ClientIP()
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
		Info("ginLog",
			String("path", path),
			String("raw", raw),
			String("mothod", mothod),
			String("clientIP", clientIP),
			Int("statusCode", statusCode),
			String("errorMessage", errorMessage))
	}

}

//String ...
func String(key string, val string) Field {
	return Field{Key: key, Type: zapcore.StringType, String: val}
}

//Int ...
func Int(key string, val int) Field {
	return Field{Key: key, Type: zapcore.Int64Type, Integer: int64(val)}
}

// Info ...
func Info(msg string, fields ...Field) {
	zap.L().Info(msg, fields...)
}

// Debug ...
func Debug(msg string, fields ...Field) {
	zap.L().Debug(msg, fields...)
}

//Error ...
func Error(msg string, fields ...Field) {
	zap.L().Error(msg, fields...)
}

//Fatal ...
func Fatal(msg string, fields ...Field) {
	zap.L().Fatal(msg, fields...)
}

//Sync ...
func Sync() {
	zap.L().Sync()
}

//ReplaceGlobals 替换logger
func ReplaceGlobals(time string) {
	cfg := zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoderCfg,
		OutputPaths:      []string{"welfare_" + time + ".log"},
		ErrorOutputPaths: []string{"welfare_" + time + ".log"},
	}
	logger, _ := cfg.Build(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}

// UpdateLogLevel 更新日志的级别
func UpdateLogLevel(s string) {
	level.UnmarshalText([]byte(s))
}
