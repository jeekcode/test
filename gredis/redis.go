package gredis

import (
	"time"
	"welfare/setting"

	"github.com/gomodule/redigo/redis"
)

var RedisConn *redis.Pool

func init() {
	RedisConn = &redis.Pool{
		MaxIdle:   10,
		MaxActive: 10,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
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
	return RedisConn.Get()
}
