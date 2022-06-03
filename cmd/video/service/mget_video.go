package service

import (
	"context"

	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetVideoByUserIdService struct {
	ctx context.Context
}

//
func NewMGetVideoByUserIdService(ctx context.Context) *MGetVideoByUserIdService {
	return &MGetVideoByUserIdService{ctx: ctx}
}

// 从DAO层获取视频基本信息，并查出点赞数、评论数、当前用户是否点赞，组装后返回
func (s *MGetVideoByUserIdService) MGetVideo(req *videoproto.GetVideosByUserIdReq) ([]*videoproto.VideoInfo, error) {
	videoModels, err := db.MGetVideoByUserId(s.ctx, req.UserId)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, err
	}

	videos := pack.Videos(videoModels) // 做类型转换：视频id、base_info已经得到，还需要点赞数、评论数、是否点赞

	// 把视频的其他信息进行绑定
	for i := 0; i < len(videos); i++ {
		vid := videos[i].VideoId
		uid := videos[i].VideoBaseInfo.UserId
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

		isFavaorite, err := db.IsFavorite(s.ctx, vid, uid)
		if err != nil {
			return nil, err
		}
		videos[i].IsFavorite = isFavaorite
	}
	return videos, nil
}
