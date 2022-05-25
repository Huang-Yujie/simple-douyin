package db

import (
	"context"
	"errors"
	videoproto2 "simple-douyin/kitex_gen/videoproto"
	"strconv"
)


// CreateVideo create video info
func CreateVideo(ctx context.Context, notes []*Video) error {
	videos := make([]Video, len(notes))
	for i:=0; i<len(videos); i++ {
		video := *notes[i]
		videos[i] = video
	}

	if err := DB.WithContext(ctx).Create(&videos).Error; err!=nil {
		return err
	}
	return nil
}

// MGetVideoByUserId 根据用户id查视频
func MGetVideoByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	var videos []*Video
	result := DB.WithContext(ctx).Model(&Video{}).Where("user_id = ?", userId)
	if result.Error != nil {
		return videos, result.Error
	}
	result.Find(&videos)
	return videos, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(ctx context.Context, vid int64) (int64, error) {
	result := DB.WithContext(ctx).Model(&Like{}).Where("video_video_id = ?", vid)
	if result.Error != nil {
		return -1, result.Error
	}
	var cnt int64
	result.Count(&cnt)
	return cnt, nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(ctx context.Context, vid int64) (int64, error) {
	var cnt int64
	result := DB.WithContext(ctx).Model(&Comment{}).Where("video_id = ?", vid)
	if result.Error != nil {
		return -1, result.Error
	}
	result.Count(&cnt)
	return cnt, nil
}

// IsFavorite 返回是否点赞
func IsFavorite(ctx context.Context, vid int64, uid int64) (bool, error) {
	result := DB.WithContext(ctx).Model(&Like{}).Where("user_user_id = ? AND video_video_id = ?", uid, vid)
	if result.Error != nil {
		return false, result.Error
	}
	var cnt int64
	result.Count(&cnt)
	return cnt>0, nil
}

// MGetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func MGetVideoByTime(ctx context.Context, LatestTime int64, Count int64) ([]*Video, int64, error) {
	var videos []*Video
	result := DB.WithContext(ctx).Model(&Video{}).Where("created_at < ?", LatestTime).Limit(int(Count)).Order("created_at").Find(&videos)
	if result.Error != nil {
		return videos, -1, result.Error
	}
	if videos == nil || len(videos) == 0 {
		return videos, -1, errors.New("可选取的视频为空")
	}
	return videos, int64(videos[0].CreatedAt), nil
}

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	like := &Like{
		UserId: uint(UserId),
		VideoId: uint(VideoId),
	}
	if err := DB.WithContext(ctx).Model(&Like{}).Create(&like).Error; err != nil {
		return err
	}
	return nil
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	result := DB.WithContext(ctx).Model(&Like{}).Where("user_user_id = ? AND video_video_id = ?", UserId, VideoId).Delete(&Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// CreateComment 新增评论,需要da层返回评论详情
func CreateComment(ctx context.Context, UserId int64, VideoId int64, Content string) (*videoproto2.CommentInfo, error) {
	comment := Comment{
		UserID: uint(UserId),
		VideoID: uint(VideoId),
		Content: Content,
	}
	if err := DB.WithContext(ctx).Model(&Comment{}).Create(&comment).Error; err != nil {
		panic(err)
	}

	resp := &videoproto2.CommentInfo{
		CommentId: int64(comment.CommentID),
		UserId: int64(comment.UserID),
		Content: comment.Content,
		CreateDate: strconv.Itoa(comment.CreatedAt),
	}

	return resp, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, CommentId int64) error {
	if err := DB.WithContext(ctx).Model(&Comment{}).Delete(&Comment{}, CommentId).Error; err != nil {
		return err
	}
	return nil
}

// GetComment 查询评论,需要da层返回评论详情,有可能有多条评论
func GetComment(ctx context.Context, UserId int64, VideoId int64) ([]*videoproto2.CommentInfo, error) {
	var comments []*Comment
	result := DB.WithContext(ctx).Model(&Comment{}).Where("user_id = ? AND video_id = ?", UserId, VideoId).Find(&comments)
	if result.Error != nil {
		return make([]*videoproto2.CommentInfo, 0), result.Error
	}
	if len(comments) == 0 {
		return make([]*videoproto2.CommentInfo, 0), errors.New("没有相应的评论信息")
	}
	commmentInfos := make([]*videoproto2.CommentInfo, len(comments))
	for i:=0; i<len(comments); i++ {
		commmentInfos[i] = &videoproto2.CommentInfo{
			CommentId: int64(comments[i].CommentID),
			UserId: int64(comments[i].UserID),
			Content: comments[i].Content,
			CreateDate: strconv.Itoa(comments[i].CreatedAt),
		}
	}

	return commmentInfos, nil
}

func CreateUser(ctx context.Context, username, encpassword string) error {
	result := DB.WithContext(ctx).Create(&User{
		Username: username,
		EncryptedPassword: encpassword,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}