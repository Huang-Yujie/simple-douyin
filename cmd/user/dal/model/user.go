package model

type User struct {
	UserID            uint   `gorm:"primarykey"`
	Username          string `gorm:"not null;unique"`
	EncryptedPassword string `gorm:"not null"`
	Videos            []Video
	Comments          []Comment
	Likes             []Video `gorm:"many2many:like"`
	Fans              []User  `gorm:"many2many:follow"`
}


func (u *User) TableName() string {
	return "user"
}
