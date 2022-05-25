package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

const (
	MYSQLDSN = "root:nhk19980516@tcp(121.41.112.176:13306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
)

// init init DB
func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQLDSN),
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
}

func GenerateTables() {
	var err error
	DB, err = gorm.Open(mysql.Open(MYSQLDSN),
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

	DB.AutoMigrate(new(User), new(Video), new(Comment))
}