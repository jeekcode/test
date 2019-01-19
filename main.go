package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jeekcode/test/logger"
	"github.com/jeekcode/test/router"
	"github.com/jeekcode/test/setting"
)

var mTime int64

func main() {

	fmt.Println("Hello")
	setting.SetUp()
	logger.Info("start")
	go Tick()
	r := router.InitRouter()
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		ReadTimeout:  5,
		WriteTimeout: 5,
	}
	server.ListenAndServe()

	// c := gredis.Get()
	// c.Do("SET", "TEST", "1")
	// v, err := c.Do("GET", "TEST")
	// value, _ := redis.String(v, err)
	// fmt.Println(value)
	// c.Close()
	// logger.Debug("replace debug")
	// logger.Info("replace info")
	// logger.Sync()
}
func Tick() {
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
			now := time.Now()
			unix := now.Unix()
			if now.Minute() == 0 && now.Hour() == 0 && unix-mTime > 300 {
				mTime = unix
				timeS := now.Format("2019-01-06")
				logger.ReplaceGlobals(timeS)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
