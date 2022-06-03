package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
)

type CreateVideoService struct {
	ctx context.Context
}

//
func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// 这个是dal层定义的
// type Video struct {
// 	UserID   int64  `json:"user_id"`
// 	Title    string `json:"title"`
// 	PlayAddr string `json:"play_addr"`
// 	CoverUrl string `json:"cover_url"`
// }

//
func (s *CreateVideoService) CreateVideo(req *videoproto.CreateVideoReq) error {
	VideoModel := &db.Video{
		UserId:   uint(req.VideoBaseInfo.UserId),
		Title:    req.VideoBaseInfo.Title,
		PlayURL:  req.VideoBaseInfo.PlayAddr,
		CoverURL: req.VideoBaseInfo.CoverAddr,
	}
	// 如果添加失败，返回error
	return db.CreateVideo(s.ctx, []*db.Video{VideoModel})
}
