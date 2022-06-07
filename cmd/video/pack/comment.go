package pack

import (
	"simple-douyin/cmd/video/dal/model"
	"simple-douyin/kitex_gen/videoproto"
	"time"
)

func Comment(comment *model.Comment) *videoproto.CommentInfo {
	return &videoproto.CommentInfo{
		CommentId:  int64(comment.CommentID),
		UserId:     int64(comment.UserID),
		Content:    comment.Content,
		CreateDate: time.Unix(int64(comment.CreatedAt), 0).Format("2006-01-02 15:04:05")[5:10],
	}
}

func Comments(comments []*model.Comment) []*videoproto.CommentInfo {
	commentInfos := make([]*videoproto.CommentInfo, len(comments))
	for i, comment := range comments {
		commentInfos[i] = Comment(comment)
	}
	return commentInfos
}
