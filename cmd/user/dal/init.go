package dal

import (
	"simple-douyin/cmd/user/dal/model"
	"simple-douyin/pkg/config"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

// init init DB
func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Database.DSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		klog.Fatal(err)
	}

	DB.AutoMigrate(new(model.User), new(model.Video), new(model.Comment))
}
