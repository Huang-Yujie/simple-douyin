package db

type Like struct {
	UserId uint `gorm:"column:user_user_id"`
	VideoId uint `gorm:"column:video_video_id"`
}

func (l *Like) TableName() string {
	return "like"
}
