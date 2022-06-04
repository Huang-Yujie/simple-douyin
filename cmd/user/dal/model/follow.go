package model

type Follow struct {
	UserId uint `gorm:"column:user_user_id"`
	FanId  uint `gorm:"column:fan_user_id"`
}

func (f *Follow) TableName() string {
	return "follow"
}
