package cache

import (
	"simple-douyin/pkg/config"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func Init() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.Server.RedisPort)
		},
	}
}
