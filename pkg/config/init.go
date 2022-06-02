package config

import (
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

func Init() {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")
	if err := vp.ReadInConfig(); err != nil {
		klog.Fatal(err)
	}
	vp.UnmarshalKey("Server", &Server)
	vp.UnmarshalKey("OSS", &OSS)
	vp.UnmarshalKey("Database", &Database)
	vp.UnmarshalKey("JWT", &JWT)
	JWT.Expires *= time.Hour
}
