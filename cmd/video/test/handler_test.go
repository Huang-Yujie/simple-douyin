package test

import (
	"context"
	"fmt"
	"simple-douyin/cmd/video/dal/db"
	videoproto2 "simple-douyin/kitex_gen/videoproto"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	db.Init()

	impl := VideoServiceImpl{}

	req := &videoproto2.CreateVideoReq{
		VideoBaseInfo: &videoproto2.VideoBaseInfo{
			UserId: int64(1),
			PlayAddr: "Testing TestCreateVideo",
			CoverAddr: "Testing TestCreateVideo",
			Title: "Testing TestCreateVideo",
		},
	}

	resp, err := impl.CreateVideo(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)

}
