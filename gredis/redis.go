package gredis

import (
	"time"

	"github.com/jeekcode/test/config"

	"github.com/gomodule/redigo/redis"
)

var redisConn *redis.Pool

func init() {
	redisConn = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 10,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
func Get() redis.Conn {
	return redisConn.Get()
}
