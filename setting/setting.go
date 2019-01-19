package setting

import (
	"welfare/logger"

	"go.uber.org/zap"

	"github.com/go-ini/ini"
)

type Server struct {
	ServerIP   string
	ServerPort int
}

var ServerSetting = &Server{}

type Redis struct {
	Host     string
	PassWord string
}

var RedisSetting = &Redis{}

type Loger struct {
	LogLevel string
	LogFile  string
}

var LoggerSetting = &Loger{}
var cfg *ini.File

func SetUp() {
	var err error
	cfg, err = ini.Load("conf\\app.ini")
	if err != nil {
		logger.Fatal("Fail to parse 'conf\\app.ini'", zap.String("Error", err.Error()))
	}
	ServerSetting.ServerIP = ServerIP()
	ServerSetting.ServerPort = ServerPort()
	RedisSetting.Host = Host()
	RedisSetting.PassWord = PassWord()
	LoggerSetting.LogLevel = LogLevel()
	LoggerSetting.LogFile = LogFile()
	logger.Info("replace Info settings")
	defer logger.Sync()
}

func ServerIP() string {
	return cfg.Section("server").Key("ServerIP").MustString("127.0.0.1")
}
func ServerPort() int {
	return cfg.Section("server").Key("ServerPort").MustInt(8080)
}
func Host() string {
	return cfg.Section("redis").Key("Host").MustString("127.0.0.1:6379")
}
func PassWord() string {
	return cfg.Section("redis").Key("PassWord").MustString("")
}
func LogLevel() string {
	return cfg.Section("logger").Key("LogLevel").MustString("info")
}
func LogFile() string {
	return cfg.Section("logger").Key("LogFile").MustString("welfare")
}
