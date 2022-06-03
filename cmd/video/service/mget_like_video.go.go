package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetLikeVideoService struct {
	ctx context.Context
}

//
func NewMGetLikeVideoService(ctx context.Context) *MGetLikeVideoService {
	return &MGetLikeVideoService{ctx: ctx}
}

// 从DAO层获取视频基本信息，并查出点赞数、评论数、当前用户是否点赞，组装后返回
func (s *MGetLikeVideoService) MGetLikeVideo(req *videoproto.GetLikeVideosReq) ([]*videoproto.VideoInfo, error) {
	videoModels, err := dal.MGetLikeVideo(s.ctx, req.AppUserId)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, err
	}

	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info已经得到，还需要点赞数、评论数、是否点赞

	// 把视频的其他信息进行绑定
	for i := 0; i < len(videos); i++ {
		vid := videos[i].VideoId
		likeCount, err := dal.GetLikeCount(s.ctx, vid)
		if err != nil {
			return nil, err
		}
		videos[i].LikeCount = likeCount

		commentCount, err := dal.GetCommentCount(s.ctx, vid)
		if err != nil {
			return nil, err
		}
		videos[i].CommentCount = commentCount
		videos[i].IsFavorite = true // 返回的视频都是已经点赞的
	}
	return videos, nil
}
