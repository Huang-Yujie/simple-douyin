package model

type Video struct {
	BaseModel
	AuthId  int32  `gorm:"column:authId"`
	Address string `gorm:"column:address;type:varchar(200)"`
	Profile string `gorm:"column:profile;type:varchar(200)"`
	Picture string `gorm:"column:picture;type:varchar(200)"`
}

func (Video) TableName() string {
	return "video"
}
