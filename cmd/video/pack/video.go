package pack

import (
	"simple-douyin/cmd/video/dal/model"
	"simple-douyin/kitex_gen/videoproto"
)

// Video pack video info : video to videoproto.VideoInfo
func Video(m *model.Video) *videoproto.VideoInfo {
	if m == nil {
		return nil
	}

	return &videoproto.VideoInfo{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:     int64(m.UserID),
			OssVideoId: m.OSSVideoID,
			Title:      m.Title,
		},
		VideoId: int64(m.VideoID),
	}
}

func Videos(ms []*model.Video) []*videoproto.VideoInfo {
	videos := make([]*videoproto.VideoInfo, len(ms))
	for i, m := range ms {
		if n := Video(m); n != nil {
			videos[i] = n
		}
	}
	return videos
}
