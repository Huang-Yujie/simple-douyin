package dal

import (
	"context"
	"fmt"
	"simple-douyin/cmd/user/dal/model"
	"simple-douyin/pkg/config"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

func testInit() {
	config.Init()
	DSN := config.Database.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB = DB.Debug()
	if err != nil {
		klog.Fatal(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		klog.Fatal(err)
	}

	DB.AutoMigrate(new(model.User), new(model.Video), new(model.Comment))
}

func TestGetUserByID(t *testing.T) {
	testInit()
	userID := int64(3)
	user, err := GetUserByID(context.Background(), userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(*user)
}

func TestGetUserByName(t *testing.T) {
	testInit()
	userName := "wyz"
	user, err := GetUserByName(context.Background(), userName)
	if err != nil {
		panic(err)
	}
	fmt.Println(*user)
}

func TestCreateUser(t *testing.T) {
	testInit()
	userName := "nhk"
	encPassword := "nhknhk"
	userID, err := CreateUser(context.Background(), userName, encPassword)
	if err != nil {
		panic(err)
	}
	fmt.Println(userID)
}

func TestFollowUser(t *testing.T) {
	testInit()
	userID := int64(3)
	fanID := int64(2)

	err := FollowUser(context.Background(), fanID, userID)
	if err != nil {
		panic(err.Error())
	}
}

func TestUnFollowUser(t *testing.T) {
	testInit()

	userID := int64(2)
	fanID := int64(1)

	err := UnFollowUser(context.Background(), fanID, userID)
	if err != nil {
		panic(err)
	}
}

func TestMGetFanUser(t *testing.T) {
	testInit()
	fanID := int64(1)
	userIDs, err := MGetFanUser(context.Background(), fanID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userIDs)
}

func TestMGetFollowUser(t *testing.T) {
	testInit()
	userID := int64(2)
	fanIDs, err := MGetFollowUser(context.Background(), userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(fanIDs)
}

func TestGetFanCount(t *testing.T) {
	testInit()
	userID := int64(2)

	count, err := GetFanCount(context.Background(), userID)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func TestGetFollowCount(t *testing.T) {
	testInit()
	fanID := int64(1)

	count, err := GetFollowCount(context.Background(), fanID)
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}

func TestIsFollow(t *testing.T) {
	testInit()

	userID := int64(2)
	fanID := int64(1)

	f, err := IsFollow(context.Background(), userID, fanID)
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}
