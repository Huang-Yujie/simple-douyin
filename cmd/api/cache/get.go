package cache

import (
	"simple-douyin/pkg/oss"

	"github.com/gomodule/redigo/redis"
)

func GetPlayURL(ossVideoID string) (string, error) {
	playURL, err := redis.String(conn.Do("GET", ossVideoID))
	// 缓存命中
	if err == nil {
		return playURL, nil
	}
	// 缓存未命中
	playURL, err = oss.GetPlayURL(ossVideoID)
	if err != nil {
		return "", err
	}
	return playURL, SetPlayURL(ossVideoID, playURL)
}
