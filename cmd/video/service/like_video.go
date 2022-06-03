package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
)

type LikeVideoService struct {
	ctx context.Context
}

//
func NewLikeVideoService(ctx context.Context) *LikeVideoService {
	return &LikeVideoService{ctx: ctx}
}

//
func (s *LikeVideoService) LikeVideo(req *videoproto.LikeVideoReq) error {

	// 如果插入错误，返回error
	return db.LikeVideo(s.ctx, req.UserId, req.VideoId)
}
