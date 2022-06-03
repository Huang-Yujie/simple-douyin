package service

import (
	"context"
	"simple-douyin/cmd/video/dal"
	"simple-douyin/kitex_gen/videoproto"
)

type UnLikeVideoService struct {
	ctx context.Context
}

func NewUnLikeVideoService(ctx context.Context) *UnLikeVideoService {
	return &UnLikeVideoService{ctx: ctx}
}

func (s *UnLikeVideoService) UnLikeVideo(req *videoproto.UnLikeVideoReq) error {
	// 如果删除错误，返回error
	return dal.UnLikeVideo(s.ctx, req.UserId, req.VideoId)
}
