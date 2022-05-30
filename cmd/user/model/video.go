package model

type Video struct {
	VideoID   uint `gorm:"primarykey"`
	UserId    int64	`gorm:"not null;comment:发表视频用户ID"`
	Title     string `gorm:"not null"`
	PlayURL   string `gorm:"not null;comment:视频源"`
	CoverURL  string `gorm:"not null;comment:视频封面"`
	CreatedAt int
	Comments  []Comment
}

func (v *Video) TableName() string {
	return "video"
}
