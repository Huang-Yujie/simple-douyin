package main

import (
	"context"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/cmd/video/service"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videoproto.CreateVideoReq) (resp *videoproto.CreateVideoResp, err error) {
	resp = new(videoproto.CreateVideoResp)

	if req.VideoBaseInfo.UserId <= 0 || len(req.VideoBaseInfo.PlayAddr) == 0 || len(req.VideoBaseInfo.CoverAddr) == 0 || len(req.VideoBaseInfo.Title) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetVideosByUserId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByUserId(ctx context.Context, req *videoproto.GetVideosByUserIdReq) (resp *videoproto.GetVideosByUserIdResp, err error) {
	resp = new(videoproto.GetVideosByUserIdResp)

	if req.UserId < 0 || req.AppUserId < 0 { // 这个id的判断规则是什么？
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	videos, err := service.NewMGetVideoByUserIdService(ctx).MGetVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoInfos = videos
	return resp, nil
}

// GetVideosByTime implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideosByTime(ctx context.Context, req *videoproto.GetVideosByTimeReq) (resp *videoproto.GetVideosByTimeResp, err error) {
	resp = new(videoproto.GetVideosByTimeResp)
	if (req.AppUserId < 0 && req.AppUserId != -1) || req.Count > 1000 || req.LatestTime < 0 { // count限制小于1000，避免查询过多数据;时间戳怎么判断？
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	videos, nextTime, err := service.NewMGetVideoByTimeService(ctx).MGetVideoByTime(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoInfos = videos
	resp.NextTime = nextTime
	return resp, nil
}

// LikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) LikeVideo(ctx context.Context, req *videoproto.LikeVideoReq) (resp *videoproto.LikeVideoResp, err error) {
	resp = new(videoproto.LikeVideoResp)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	err = service.NewLikeVideoService(ctx).LikeVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UnLikeVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) UnLikeVideo(ctx context.Context, req *videoproto.UnLikeVideoReq) (resp *videoproto.UnLikeVideoResp, err error) {
	resp = new(videoproto.UnLikeVideoResp)

	if req.UserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	err = service.NewUnLikeVideoService(ctx).UnLikeVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CreateComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateComment(ctx context.Context, req *videoproto.CreateCommentReq) (resp *videoproto.CreateCommentResp, err error) {
	resp = new(videoproto.CreateCommentResp)

	if req.UserId <= 0 || req.VideoId <= 0 || len(req.Content) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	commentInfo, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentInfo = commentInfo
	return resp, nil
}

// DeleteComment implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DeleteComment(ctx context.Context, req *videoproto.DeleteCommentReq) (resp *videoproto.DeleteCommentResp, err error) {
	resp = new(videoproto.DeleteCommentResp)

	if req.CommentId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetComments implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetComments(ctx context.Context, req *videoproto.GetCommentsReq) (resp *videoproto.GetCommentsResp, err error) {
	resp = new(videoproto.GetCommentsResp)

	if req.AppUserId < 0 || req.VideoId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	comments, err := service.NewMGetCommentService(ctx).MGetComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentInfos = comments
	return resp, nil
}
