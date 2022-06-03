package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/kitex_gen/videoproto"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *videoproto.CreateCommentReq) (*videoproto.CommentInfo, error) {

	// 新增评论，需要DAO层返回评论详情
	return dal.CreateComment(s.ctx, req.UserId, req.VideoId, req.Content)
}
