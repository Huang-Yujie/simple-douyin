package cache

import "simple-douyin/pkg/config"

func SetPlayURL(ossVideoID string, playURL string) error {
	_, err := conn.Do("SET", ossVideoID, playURL)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", ossVideoID, config.Server.Timeout)
	return err
}
