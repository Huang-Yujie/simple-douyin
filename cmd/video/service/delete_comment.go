package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/kitex_gen/videoproto"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteComment(req *videoproto.DeleteCommentReq) error {
	return dal.DeleteComment(s.ctx, req.CommentId)
}
