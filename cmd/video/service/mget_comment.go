package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetCommentService struct {
	ctx context.Context
}

func NewMGetCommentService(ctx context.Context) *MGetCommentService {
	return &MGetCommentService{ctx: ctx}
}

func (s *MGetCommentService) MGetComment(req *videoproto.GetCommentsReq) ([]*videoproto.CommentInfo, error) {
	// 查询评论,需要dal层返回评论详情,有可能有多条评论
	comments, err := dal.MGetComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return pack.Comments(comments), nil
}
