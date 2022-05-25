package db

import "context"

type Comment struct {
	CommentID uint `gorm:"primarykey"`
	UserID    uint
	VideoID   uint
	Content   string `gorm:"not null"`
	CreatedAt int
}

func NewCommentDB() *Comment {
	return &Comment{}
}

func (c *Comment) TableName() string {
	return "comment"
}

func(c *Comment) InsertOneRecord(ctx context.Context, userId uint, videoId uint, content string) error {
	comment := Comment{
		UserID: userId,
		VideoID: videoId,
		Content: content,
	}
	if err := DB.WithContext(ctx).Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func(c *Comment) GetCommentIdByUserIdAndVideo(ctx context.Context, userId uint, videoId uint) ([]*Comment, error) {
	var comments []*Comment
	result := DB.WithContext(ctx).Model(&Comment{}).Where("user_id = ? AND video_id = ?", userId, videoId).Find(&comments)
	if result.Error != nil {
		return make([]*Comment, 0), result.Error
	}
	return comments, nil
}

func(c *Comment) UpdateContent(ctx context.Context, commentId uint, content string) error {
	return nil
}

func(c *Comment) SelectCommentByUserId(ctx context.Context, userId uint, videoId uint) ([]*Comment, error) {
	return nil, nil
}

func(c *Comment) SelectCommentByVideoId(ctx context.Context, userId uint, videoId uint) ([]*Comment, error) {
	return nil, nil
}





