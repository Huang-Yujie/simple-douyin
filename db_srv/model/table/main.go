package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"simple-douyin/db_srv/model"
)

func main() {
	db,err := gorm.Open(mysql.Open("root:nhk19980516@(121.41.112.176:13306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(&model.Video{}); err != nil {
		panic(err)
	}
}

