package test

import (
	"context"
	"simple-douyin/cmd/video/dal/db"
	videoproto2 "simple-douyin/kitex_gen/videoproto"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto2.CreateVideoReq) (resp *videoproto2.CreateVideoResp, err error) {
	// TODO: Your code here...
	userId := uint(req.VideoBaseInfo.UserId)
	playAddr := req.VideoBaseInfo.PlayAddr
	coverAddr := req.VideoBaseInfo.CoverAddr
	title := req.VideoBaseInfo.Title

	videoDB := db.NewVideoDB()
	if err = videoDB.InsertOneRecord(ctx, userId, playAddr, coverAddr, title); err != nil {
		return nil, err
	}

	resp = &videoproto2.CreateVideoResp{
		BaseResp: &videoproto2.BaseResp{
			StatusCode: 200,
			StatusMsg: "video: 插入成功",
		},
	}
	return resp, nil
}

// GetVideosByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByUserId(ctx context.Context, req *videoproto2.GetVideosByUserIdReq) (resp *videoproto2.GetVideosByUserIdResp, err error) {
	// TODO: Your code here...
	return
}

// GetVideosByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByTime(ctx context.Context, req *videoproto2.GetVideosByTimeReq) (resp *videoproto2.GetVideosByTimeResp, err error) {
	// TODO: Your code here...
	return
}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *videoproto2.LikeVideoReq) (resp *videoproto2.LikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// UnLikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLikeVideo(ctx context.Context, req *videoproto2.UnLikeVideoReq) (resp *videoproto2.UnLikeVideoResp, err error) {
	// TODO: Your code here...
	return
}

// CreateComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateComment(ctx context.Context, req *videoproto2.CreateCommentReq) (resp *videoproto2.CreateCommentResp, err error) {
	// TODO: Your code here...
	return
}

// DeleteComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteComment(ctx context.Context, req *videoproto2.DeleteCommentReq) (resp *videoproto2.DeleteCommentResp, err error) {
	// TODO: Your code here...
	return
}

// GetComments implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetComments(ctx context.Context, req *videoproto2.GetCommentsReq) (resp *videoproto2.GetCommentsResp, err error) {
	// TODO: Your code here...
	return
}
