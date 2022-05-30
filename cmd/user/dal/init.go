package dal

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"simple-douyin/cmd/user/model"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

const (
	MYSQLDSN = "root:123456@tcp(172.31.68.40:9910)/"
	DBName = "douyin"
)

// 初始化数据库创建函数
func initializeMysqlDB() {
	// 连接MySQL的服务器
	db, err := sql.Open("mysql", MYSQLDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// 构建SQL语句
	sqlStr := "CREATE DATABASE IF NOT EXISTS "+ DBName + " Character Set utf8mb4"
	// 创建数据库
	if _, err = db.Exec(sqlStr); err != nil {
		panic("数据库创建失败！")
	}
}

// Init init DB
func Init() {
	initializeMysqlDB()
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQLDSN+DBName),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB.AutoMigrate(new(model.User), new(model.Video), new(model.Comment))
}
