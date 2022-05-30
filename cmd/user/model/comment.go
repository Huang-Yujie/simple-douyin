package model


type Comment struct {
	CommentID uint 	`gorm:"primarykey;comment:评论ID"`
	UserID    int64	`gorm:"comment:评论者ID"`
	VideoID   uint	`gorm:"comment:视频ID"`
	Content   string `gorm:"not null;comment:评论内容"`
	CreatedAt int	`gorm:"comment:评论创建时间"`
}

func (c *Comment) TableName() string {
	return "comment"
}
