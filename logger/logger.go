package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zapcore.Field
type HandelFunc = func(*gin.Context)

var encoder_cfg = zapcore.EncoderConfig{
	TimeKey:        "Time",
	LevelKey:       "Level",
	NameKey:        "Name",
	CallerKey:      "Call",
	MessageKey:     "Msg",
	StacktraceKey:  "S",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.CapitalLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.StringDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}
var level = zap.NewAtomicLevel()

func init() {
	timeS := time.Now().Format("2006-01-02")
	cfg := zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"welfare" + timeS + ".log"},
		ErrorOutputPaths: []string{"welfare" + timeS + ".log"},
	}
	logger, err := cfg.Build(zap.AddCallerSkip(1))
	if err != nil {
		zap.L().Fatal("logger init error")
	}
	zap.ReplaceGlobals(logger)
}

func LogGin() HandelFunc {
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
		zap.L().Info("ginLog",
			zap.String("path", path),
			zap.String("raw", raw),
			zap.String("mothod", mothod),
			zap.String("clientIP", clientIP),
			zap.Int("statusCode", statusCode),
			zap.String("errorMessage", errorMessage))
	}

}
func Info(msg string, fields ...Field) {
	zap.L().Info(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	zap.L().Debug(msg, fields...)
}

func Error(msg string, fields ...Field) {
	zap.L().Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	zap.L().Fatal(msg, fields...)
}

func Sync() {
	zap.L().Sync()
}

func ReplaceGlobals(time string) {
	cfg := zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoder_cfg,
		OutputPaths:      []string{"welfare_" + time + ".log"},
		ErrorOutputPaths: []string{"welfare_" + time + ".log"},
	}
	logger, _ := cfg.Build(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(logger)
}
func UpdateLogLevel(s string) {
	level.UnmarshalText([]byte(s))
}
