package pack

import (
	"simple-douyin/cmd/video/dal/db"
	"simple-douyin/kitex_gen/videoproto"
)

// Video pack video info : video to videoproto.VideoInfo
func Video(m *db.Video) *videoproto.VideoInfo {
	if m == nil {
		return nil
	}

	return &videoproto.VideoInfo{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:    int64(m.UserId),
			PlayAddr:  m.PlayURL,
			CoverAddr: m.CoverURL,
			Title:     m.Title,
		},
		VideoId: int64(m.VideoID),
	}
}

func Videos(ms []*db.Video) []*videoproto.VideoInfo {
	videos := make([]*videoproto.VideoInfo, 0)
	for _, m := range ms {
		if n := Video(m); n != nil {
			videos = append(videos, n)
		}
	}
	return videos
}
