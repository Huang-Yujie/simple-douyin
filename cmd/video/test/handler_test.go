package test

import (
	"context"
	"fmt"
	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/cmd/video/kitex_gen/videoproto"
	"testing"
)

func TestCreateVideo(t *testing.T) {
	db.Init()

	impl := VideoServiceImpl{}

	req := &videoproto.CreateVideoReq{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
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
