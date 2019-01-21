package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jeekcode/test/config"
	"github.com/jeekcode/test/logger"
	"github.com/jeekcode/test/router"
)

var mTime int64

func main() {
	logger.Init()
	fmt.Println("Hello")
	config.SetUp()
	go Tick()
	r := router.InitRouter()
	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		ReadTimeout:  5,
		WriteTimeout: 5,
	}
	server.ListenAndServe()
}

//Tick ...
func Tick() {
	tick := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-tick.C:
			now := time.Now()
			unix := now.Unix()
			if now.Minute() == 0 && now.Hour() == 0 && unix-mTime > 300 {
				mTime = unix
				timeS := now.Format("2006-01-02")
				logger.ReplaceGlobals(timeS)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
