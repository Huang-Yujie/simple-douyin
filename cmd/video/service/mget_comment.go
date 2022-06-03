package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetCommentService struct {
	ctx context.Context
}

//
func NewMGetCommentService(ctx context.Context) *MGetCommentService {
	return &MGetCommentService{ctx: ctx}
}

//
func (s *MGetCommentService) MGetComment(req *videoproto.GetCommentsReq) ([]*videoproto.CommentInfo, error) {

	// 查询评论,需要da层返回评论详情,有可能有多条评论
	return db.GetComment(s.ctx, req.AppUserId, req.VideoId)
}
