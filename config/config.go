package config

import (
	"os"

	"github.com/jeekcode/test/logger"

	"github.com/go-ini/ini"
)

type server struct {
	ServerIP   string
	ServerPort int
}

//ServerSetting ...
var ServerSetting = &server{}

type redis struct {
	Host     string
	PassWord string
}

//RedisSetting ...
var RedisSetting = &redis{}

type log struct {
	LogLevel string
	LogFile  string
}

//LoggerSetting ...
var LoggerSetting = &log{}
var cfg *ini.File

//SetUp init
func SetUp() {
	var err error
	cfg, err = ini.Load("conf" + string(os.PathSeparator) + "app.ini")
	if err != nil {
		logger.Fatal("Fail to parse 'conf\\app.ini'", logger.String("Error", err.Error()))
	}
	ServerSetting.ServerIP = ServerIP()
	ServerSetting.ServerPort = ServerPort()
	RedisSetting.Host = Host()
	RedisSetting.PassWord = PassWord()
	LoggerSetting.LogLevel = LogLevel()
	logger.UpdateLogLevel(LoggerSetting.LogFile)
}

//ServerIP 服务器的IP
func ServerIP() string {
	return cfg.Section("server").Key("ServerIP").MustString("127.0.0.1")
}

//ServerPort 服务器的端口
func ServerPort() int {
	return cfg.Section("server").Key("ServerPort").MustInt(8080)
}

//Host redis host
func Host() string {
	return cfg.Section("redis").Key("Host").MustString("127.0.0.1:6379")
}

//PassWord redis passwd
func PassWord() string {
	return cfg.Section("redis").Key("PassWord").MustString("")
}

//LogLevel log level
func LogLevel() string {
	return cfg.Section("logger").Key("LogLevel").MustString("info")
}
