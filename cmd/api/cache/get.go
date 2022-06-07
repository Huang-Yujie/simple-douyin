package cache

import (
	"simple-douyin/pkg/oss"

	"github.com/gomodule/redigo/redis"
)

func GetPlayURL(ossVideoID string) (string, error) {
	c := pool.Get()
	defer c.Close()
	playURL, err := redis.String(c.Do("GET", ossVideoID))
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

func GetCoverURL(ossVideoID string) (string, error) {
	c := pool.Get()
	defer c.Close()
	coverKey := ossVideoID + "cover"
	coverURL, err := redis.String(c.Do("GET", coverKey))
	// 缓存命中
	if err == nil {
		return coverURL, nil
	}
	// 缓存未命中
	coverURL, err = oss.GetCoverURL(ossVideoID)
	if err != nil {
		return "", err
	}
	return coverURL, SetCoverURL(coverKey, coverURL)
}
