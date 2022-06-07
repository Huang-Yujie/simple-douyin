package service

import (
	"context"

	"simple-douyin/cmd/video/dal"
	"simple-douyin/cmd/video/dal/model"
	"simple-douyin/kitex_gen/videoproto"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

// 这个是dal层定义的
// type Video struct {
// 	UserID   int64  `json:"user_id"`
// 	Title    string `json:"title"`
// 	OSSVideoID string `json:"oss_video_id"`
// }

func (s *CreateVideoService) CreateVideo(req *videoproto.CreateVideoReq) error {
	video := &model.Video{
		UserID:     uint(req.VideoBaseInfo.UserId),
		Title:      req.VideoBaseInfo.Title,
		OSSVideoID: req.VideoBaseInfo.OssVideoId,
	}
	// 如果添加失败，返回error
	return dal.CreateVideo(s.ctx, video)
}
