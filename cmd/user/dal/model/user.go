package model

type User struct {
	UserID            int64   `gorm:"primarykey"`
	UserName          string `gorm:"not null;unique;comment:用户名"`
	EncryptedPassword string `gorm:"not null;comment:用户密码"`
	Videos            []Video
	Likes             []Video `gorm:"many2many:like"`
	Fans              []User  `gorm:"many2many:follow"`
}

func (u *User) TableName() string {
	return "user"
}
