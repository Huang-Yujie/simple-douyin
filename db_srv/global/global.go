package global

import (
	"gorm.io/gorm"
	"simple-douyin/db_srv/config"
)

var (
	DB *gorm.DB
	DBConfig *config.MysqlConfig
)
