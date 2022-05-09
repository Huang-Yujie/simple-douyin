package model

type User struct {
	BaseModel
	Mobile   string `gorm:"column:mobile;index:idx_mobile;unique;type:varchar(11);not null"`
	Password string `gorm:"column:password;type:varchar(100);not null"`
	NickName string `gorm:"column:nickname;type:varchar(20)"`
	Gender   int    `gorm:"column:gender;default:0;type:int comment '0表示女1表示男'"`
	Picture  string `gorm:"column:picture;type:varchar(100) comment '表示用户头像图片存储位置'"`
	Profile  string `gorm:"column:profile;varchar(200) comment '表示用户个人简介'"`
}

func (User) TableName() string {
	return "user"
}