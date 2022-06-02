package cache

import (
	"simple-douyin/pkg/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gomodule/redigo/redis"
)

var conn redis.Conn

func Init() {
	var err error
	conn, err = redis.Dial("tcp", config.Server.RedisPort)
	if err != nil {
		klog.Fatal(err)
	}
}
