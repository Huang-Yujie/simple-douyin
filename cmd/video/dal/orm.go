package dal

import (
	"context"
	"errors"
	"simple-douyin/cmd/video/dal/model"
	"simple-douyin/kitex_gen/videoproto"
	"strconv"
	"time"
)

// CreateVideo create video info
func CreateVideo(ctx context.Context, video *model.Video) error {
	if err := DB.WithContext(ctx).Create(&video).Error; err != nil {
		return err
	}
	return nil
}

// MGetVideoByUserId 根据用户id查视频
func MGetVideoByUserId(ctx context.Context, userId int64) ([]*model.Video, error) {
	var videos []*model.Video
	result := DB.WithContext(ctx).Where("user_id = ?", userId)
	if result.Error != nil {
		return videos, result.Error
	}
	result.Find(&videos)
	return videos, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(ctx context.Context, vid int64) (int64, error) {
	result := DB.WithContext(ctx).Where("video_video_id = ?", vid)
	if result.Error != nil {
		return -1, result.Error
	}
	var cnt int64
	result.Count(&cnt)
	return cnt, nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(ctx context.Context, vid int64) (int64, error) {
	result := DB.WithContext(ctx).Where("video_id = ?", vid)
	if result.Error != nil {
		return -1, result.Error
	}
	var cnt int64
	result.Count(&cnt)
	return cnt, nil
}

// IsFavorite 返回是否点赞
func IsFavorite(ctx context.Context, vid int64, uid int64) (bool, error) {
	result := DB.WithContext(ctx).Where("user_user_id = ? AND video_video_id = ?", uid, vid)
	if result.Error != nil {
		return false, result.Error
	}
	var cnt int64
	result.Count(&cnt)
	return cnt > 0, nil
}

// MGetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func MGetVideoByTime(ctx context.Context, LatestTime int64, Count int64) ([]*model.Video, int64, error) {
	var videos []*model.Video
	result := DB.WithContext(ctx).Where("created_at < ?", LatestTime).Limit(int(Count)).Order("created_at").Find(&videos)
	if result.Error != nil {
		return videos, -1, result.Error
	}
	if len(videos) == 0 {
		return videos, -1, errors.New("可选取的视频为空")
	}
	return videos, int64(videos[0].CreatedAt), nil
}

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	like := &model.Like{
		UserId:  uint(UserId),
		VideoId: uint(VideoId),
	}
	if err := DB.WithContext(ctx).Create(&like).Error; err != nil {
		return err
	}
	return nil
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(ctx context.Context, UserId int64, VideoId int64) error {
	result := DB.WithContext(ctx).Where("user_user_id = ? AND video_video_id = ?", UserId, VideoId).Delete(&model.Like{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func MGetLikeVideo(ctx context.Context, userId int64) ([]*model.Video, error) {
	var videos []*model.Video
	result := DB.WithContext(ctx).Where("user_id = ?", userId).Find(&videos)
	if result.Error != nil {
		return videos, result.Error
	}
	if len(videos) == 0 {
		return videos, errors.New("可选取的视频为空")
	}
	return videos, nil
}

// CreateComment 新增评论,需要dal层返回评论详情
func CreateComment(ctx context.Context, UserId int64, VideoId int64, Content string) (*videoproto.CommentInfo, error) {
	comment := model.Comment{
		UserID:  uint(UserId),
		VideoID: uint(VideoId),
		Content: Content,
	}
	if err := DB.WithContext(ctx).Create(&comment).Error; err != nil {
		panic(err)
	}

	resp := &videoproto.CommentInfo{
		CommentId:  int64(comment.CommentID),
		UserId:     int64(comment.UserID),
		Content:    comment.Content,
		CreateDate: strconv.Itoa(comment.CreatedAt),
	}

	return resp, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, CommentId int64) error {
	if err := DB.WithContext(ctx).Delete(&model.Comment{}, CommentId).Error; err != nil {
		return err
	}
	return nil
}

// MGetComment 查询评论,需要da层返回评论详情,有可能有多条评论
func MGetComment(ctx context.Context, UserId int64, VideoId int64) ([]*videoproto.CommentInfo, error) {
	var comments []*model.Comment
	result := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", UserId, VideoId).Find(&comments)
	if result.Error != nil {
		return make([]*videoproto.CommentInfo, 0), result.Error
	}
	if len(comments) == 0 {
		return make([]*videoproto.CommentInfo, 0), errors.New("没有相应的评论信息")
	}
	commmentInfos := make([]*videoproto.CommentInfo, len(comments))
	for i := 0; i < len(comments); i++ {
		commmentInfos[i] = &videoproto.CommentInfo{
			CommentId:  int64(comments[i].CommentID),
			UserId:     int64(comments[i].UserID),
			Content:    comments[i].Content,
			CreateDate: time.Unix(int64(comments[i].CreatedAt), 0).Format("2006-01-02 15:04:05")[5:10],
		}
	}
	return commmentInfos, nil
}

func CreateUser(ctx context.Context, username, encpassword string) error {
	result := DB.WithContext(ctx).Create(&model.User{
		Username:          username,
		EncryptedPassword: encpassword,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
