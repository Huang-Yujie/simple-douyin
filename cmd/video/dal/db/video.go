package db

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



