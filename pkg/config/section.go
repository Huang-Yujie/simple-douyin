package config

import (
	"fmt"
	"time"
)

type ServerConfig struct {
	RunMode          string
	HttpPort         string
	RedisPort        string
	UserServiceAddr  string
	VideoServiceAddr string
	Timeout          int
	EtcdAddress      string
	FeedCount        int64
}

type OSSConfig struct {
	KeyID     string
	KeySecret string
	Endpoint  string
}

type DatabaseConfig struct {
	DBType    string
	UserName  string
	Password  string
	Host      string
	DBName    string
	Charset   string
	ParseTime string
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		d.UserName,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)
}

type JWTConfig struct {
	Secret  string
	Expires time.Duration
}
