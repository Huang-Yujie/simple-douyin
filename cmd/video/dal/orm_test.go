package dal

import (
	"context"
	"fmt"
	"simple-douyin/cmd/video/dal/model"
	"testing"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

func testInit() {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"),
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

func TestCreateUser(t *testing.T) {
	testInit()
	username := "kkk"
	password := "123"
	err := CreateUser(context.Background(), username, password)
	if err != nil {
		panic(err)
	}
}

func TestCreateVideo(t *testing.T) {
	testInit()
	video := &model.Video{
		UserId:   uint(1),
		Title:    "TestingCreateVideo",
		PlayURL:  "TestingCreateVideo",
		CoverURL: "TestingCreateVideo",
	}

	err := CreateVideo(context.Background(), video)
	if err != nil {
		panic(err)
	}
}

func TestMGetVideoByUserId(t *testing.T) {
	testInit()
	userId := int64(1)
	videoInfo, err := MGetVideoByUserId(context.Background(), userId)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(videoInfo); i++ {
		fmt.Println(*videoInfo[i])
	}
}

func TestGetLikeCount(t *testing.T) {
	testInit()
	// 需要先验证是否已点赞，如果点过赞不应执行插入操作
	vid := int64(3)
	cnt, err := GetLikeCount(context.Background(), vid)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

func TestGetCommentCount(t *testing.T) {
	testInit()
	vid := int64(3)
	cnt, err := GetCommentCount(context.Background(), vid)
	if err != nil {
		panic(err)
	}
	fmt.Println(cnt)
}

func TestIsFavorite(t *testing.T) {
	testInit()
	uid := int64(2)
	vid := int64(3)
	flag, err := IsFavorite(context.Background(), vid, uid)
	if err != nil {
		panic(err)
	}
	fmt.Println(flag)
}

func TestMGetVideoByTime(t *testing.T) {
	testInit()
	lastTime := time.Now().Unix()
	videos, nextTime, err := MGetVideoByTime(context.Background(), lastTime, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println(nextTime)
	for i := 0; i < len(videos); i++ {
		fmt.Println(*videos[i])
	}
}

func TestLikeVideo(t *testing.T) {
	testInit()
	userId := int64(1)
	videoId := int64(1)
	if err := LikeVideo(context.Background(), userId, videoId); err != nil {
		panic(err)
	}
}

func TestUnLikeVideo(t *testing.T) {
	testInit()
	userId := int64(3)
	videoId := int64(3)
	if err := UnLikeVideo(context.Background(), userId, videoId); err != nil {
		panic(err)
	}
}

func TestCreateComment(t *testing.T) {
	testInit()
	userId := int64(1)
	videoId := int64(1)
	content := "TestingCreateComment"
	commentInfo, err := CreateComment(context.Background(), userId, videoId, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(*commentInfo)
}

func TestDeleteComment(t *testing.T) {
	testInit()
	commentId := int64(1)
	if err := DeleteComment(context.Background(), commentId); err != nil {
		panic(err)
	}
}

func TestGetComment(t *testing.T) {
	testInit()
	userId := int64(1)
	videoId := int64(1)
	commentInfos, err := MGetComment(context.Background(), userId, videoId)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(commentInfos); i++ {
		fmt.Printf("%#v\n", *commentInfos[i])
	}
}
