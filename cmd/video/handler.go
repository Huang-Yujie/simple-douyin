package main

import (
	"context"
	"simple-douyin/kitex_gen/videoproto"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) (resp *videoproto.CreateVideoResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideosByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByUserId(ctx context.Context, req *videoproto.GetVideosByUserIdReq) (resp *videoproto.GetVideosByUserIdResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideosByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByTime(ctx context.Context, req *videoproto.GetVideosByTimeReq) (resp *videoproto.GetVideosByTimeResp, err error) {
	// TODO: Your code here...
	return
}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) (resp *videoproto.LikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// UnLikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) (resp *videoproto.UnLikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// GetWhetherBeLiked implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetWhetherBeLiked(ctx context.Context, req *videoproto.GetWhetherBeLikedReq) (resp *videoproto.GetWhetherBeLikedResp, err error) {
	// TODO: Your code here...
	return
}

// GetLikesCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetLikesCount(ctx context.Context, req *videoproto.GetLikesCountReq) (resp *videoproto.GetLikesCountResp, err error) {
	// TODO: Your code here...
	return
}

// CreateComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateComment(ctx context.Context, req *videoproto.CreateCommentReq) (resp *videoproto.CreateCommentResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteComment(ctx context.Context, req *videoproto.DeleteCommentReq) (resp *videoproto.DeleteCommentResp, err error) {
	// TODO: Your code here...
	return
}

// GetComments implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetComments(ctx context.Context, req *videoproto.GetCommentsReq) (resp *videoproto.GetCommentsResp, err error) {
	// TODO: Your code here...
	return
}

// GetCommentsCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetCommentsCount(ctx context.Context, req *videoproto.GetCommentsCountReq) (resp *videoproto.GetCommentsCountResp, err error) {
	// TODO: Your code here...
	return
}
