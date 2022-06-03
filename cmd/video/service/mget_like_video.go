package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetLikeVideosService struct {
	ctx context.Context
}

//
func NewMGetLikeVideosService(ctx context.Context) *MGetLikeVideosService {
	return &MGetLikeVideosService{ctx: ctx}
}

// 从DAO层获取视频基本信息，并查出点赞数、评论数、当前用户是否点赞，组装后返回
func (s *MGetLikeVideosService) MGetLikeVideos(req *videoproto.GetLikeVideosReq) ([]*videoproto.VideoInfo, error) {
	videoModels, err := db.MGetLikeVideos(s.ctx, req.AppUserId)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, err
	}

	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info已经得到，还需要点赞数、评论数、是否点赞

	// 把视频的其他信息进行绑定
	for i := 0; i < len(videos); i++ {
		vid := videos[i].VideoId
		likeCount, err := db.GetLikeCount(s.ctx, vid)
		if err != nil {
			return nil, err
		}
		videos[i].LikeCount = likeCount

		commentCount, err := db.GetCommentCount(s.ctx, vid)
		if err != nil {
			return nil, err
		}
		videos[i].CommentCount = commentCount
		videos[i].IsFavorite = true // 返回的视频都是已经点赞的
	}
	return videos, nil
}
