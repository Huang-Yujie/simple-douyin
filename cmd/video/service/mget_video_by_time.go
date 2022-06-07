package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/cmd/video/pack"
	"simple-douyin/kitex_gen/videoproto"
)

type MGetVideoByTimeService struct {
	ctx context.Context
}

func NewMGetVideoByTimeService(ctx context.Context) *MGetVideoByTimeService {
	return &MGetVideoByTimeService{ctx: ctx}
}

// 从DAO层获取视频基本信息，并查出点赞数、评论数、当前用户是否点赞，组装后返回
func (s *MGetVideoByTimeService) MGetVideoByTime(req *videoproto.GetVideosByTimeReq) ([]*videoproto.VideoInfo, int64, error) {
	videoModels, nextTime, err := dal.MGetVideoByTime(s.ctx, req.LatestTime, req.Count)
	// 只能得到视频id，uid，title，play_url,cover_url,created_time
	if err != nil {
		return nil, 0, err
	}
	videos := pack.Videos(videoModels) // 类型转换：视频id、base_info已经得到，还需要点赞数、评论数、是否点赞
	appUserID := req.AppUserId
	// 把视频的其他信息进行绑定
	for i := 0; i < len(videos); i++ {
		vid := videos[i].VideoId
		likeCount, err := dal.GetLikeCount(s.ctx, vid)
		if err != nil {
			return nil, 0, err
		}
		videos[i].LikeCount = likeCount

		commentCount, err := dal.GetCommentCount(s.ctx, vid)
		if err != nil {
			return nil, 0, err
		}
		videos[i].CommentCount = commentCount
		if appUserID > 0 { // 判断是否进行了登陆
			isFavorite, err := dal.IsFavorite(s.ctx, vid, appUserID)
			if err != nil {
				return nil, 0, err
			}
			videos[i].IsFavorite = isFavorite
		} else { // 如果没有登陆，则点赞直接返回false
			videos[i].IsFavorite = false
		}
	}
	return videos, nextTime, nil
}
