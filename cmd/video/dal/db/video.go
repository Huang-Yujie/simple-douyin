package db

import (
	"context"
	"errors"
	"fmt"
)

type Video struct {
	VideoID   uint   `gorm:"primarykey"`
	UserId    uint
	Title     string `gorm:"not null"`
	PlayURL   string `gorm:"not null"`
	CoverURL  string `gorm:"not null"`
	CreatedAt int
	Comments  []Comment
}

func (v *Video) TableName() string {
	return "video"
}

func NewVideoDB() *Video {
	return &Video{}
}

// InsertOneRecord 为video数据表添加一条数据
func (v *Video) InsertOneRecord(ctx context.Context, userId uint, title, playURL, coverURL string) error {
	video := Video{
		UserId: userId,
		Title: title,
		PlayURL: playURL,
		CoverURL: coverURL,
	}

	if result := DB.WithContext(ctx).Create(&video); result.Error != nil {
		return result.Error
	}
	return nil
}

// InsertRecords 为video数据表添加多条数据,
func (v *Video) InsertRecords(ctx context.Context, videoInfo map[string][]interface{}) error {
	var exist bool
	var userIdArr, titleArr, playURLArr, coverURLArr []interface{}
	if userIdArr, exist = videoInfo["userId"]; !exist {
		return errors.New(fmt.Sprintln("Video表批量插入错误，缺少必要的属性`userId`"))
	}
	if titleArr, exist = videoInfo["title"]; !exist {
		return errors.New(fmt.Sprintln("Video表批量插入错误，缺少必要的属性`title`"))
	}
	if playURLArr, exist = videoInfo["playURL"]; !exist {
		return errors.New(fmt.Sprintln("Video表批量插入错误，缺少必要的属性`playURL`"))
	}
	if coverURLArr, exist = videoInfo["coverURL"]; !exist {
		return errors.New(fmt.Sprintln("Video表批量插入错误，缺少必要的属性`coverURL`"))
	}

	videos := make([]Video, len(userIdArr))

	for i:=0; i<len(userIdArr); i++ {
		userId := uint(userIdArr[i].(int))
		title := titleArr[i].(string)
		playURL := playURLArr[i].(string)
		coverURL := coverURLArr[i].(string)
		videos[i] = Video{
			UserId: userId,
			Title: title,
			PlayURL: playURL,
			CoverURL: coverURL,
		}
	}

	if result := DB.WithContext(ctx).Create(&videos); result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateTitle 更新标题
func (v *Video) UpdateTitle(ctx context.Context, videoId uint, title string) error {
	video := Video{
		VideoID: videoId,
	}
	if err := DB.WithContext(ctx).Model(&video).Update("title", title).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePlayURL 更新播放地址
func (v *Video) UpdatePlayURL(ctx context.Context, videoId uint, playURL string) error {
	video := Video{
		VideoID: videoId,
	}
	if err := DB.WithContext(ctx).Model(&video).Update("play_url", playURL).Error; err != nil {
		return err
	}
	return nil
}

// UpdateCoverURL 更新封面地址
func (v *Video) UpdateCoverURL(ctx context.Context, videoId uint, coverURL string) error {
	video := Video{
		VideoID: videoId,
	}
	if err := DB.WithContext(ctx).Model(&video).Update("cover_url", coverURL).Error; err != nil {
		return err
	}
	return nil
}

// GetVideoByUserId 根据用户Id获取视频
func (v *Video) GetVideoByUserId(ctx context.Context, userId uint) ([]*Video, error) {
	var videos []*Video
	result := DB.WithContext(ctx).Model(&Video{}).Where("user_id = ?", userId)
	if result.Error != nil {
		return videos, result.Error
	}

	result.Find(&videos)
	return videos, nil
}



