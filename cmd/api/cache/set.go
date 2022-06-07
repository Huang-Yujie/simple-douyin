package cache

import "simple-douyin/pkg/config"

func SetPlayURL(ossVideoID string, playURL string) error {
	c := pool.Get()
	defer c.Close()
	_, err := c.Do("SET", ossVideoID, playURL)
	if err != nil {
		return err
	}
	_, err = c.Do("EXPIRE", ossVideoID, config.Server.Timeout)
	return err
}

func SetCoverURL(ossVideoID string, coverURL string) error {
	c := pool.Get()
	defer c.Close()
	coverKey := ossVideoID + "cover"
	_, err := c.Do("SET", coverKey, coverURL)
	if err != nil {
		return err
	}
	_, err = c.Do("EXPIRE", coverKey, config.Server.Timeout)
	return err

}
