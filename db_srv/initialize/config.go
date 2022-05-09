package initialize

import (
	"github.com/spf13/viper"
	"simple-douyin/db_srv/global"
)

// InitConfig 初始化配置一般初始化global内的配置结构体
func InitConfig() {
	v := viper.New()
	v.SetConfigFile("db_srv/config-debug.yaml")

	// 解析配置文件到viper中
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	// 解析viper结构体到global内
	if err := v.Unmarshal(&global.DBConfig); err != nil {
		panic(err)
	}
}