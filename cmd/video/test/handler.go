package test

import (
	"context"
	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/cmd/video/kitex_gen/videoproto"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) (resp *videoproto.CreateVideoResp, err error) {
	// TODO: Your code here...
	userId := uint(req.VideoBaseInfo.UserId)
	playAddr := req.VideoBaseInfo.PlayAddr
	coverAddr := req.VideoBaseInfo.CoverAddr
	title := req.VideoBaseInfo.Title

	videoDB := db.NewVideoDB()
	if err = videoDB.InsertOneRecord(ctx, userId, playAddr, coverAddr, title); err != nil {
		return nil, err
	}

	resp = &videoproto.CreateVideoResp{
		BaseResp: &videoproto.BaseResp{
			StatusCode: 200,
			StatusMsg: "video: 插入成功",
		},
	}
	return resp, nil
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
