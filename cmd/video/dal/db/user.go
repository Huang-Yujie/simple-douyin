package db

import "context"

type User struct {
	UserID            uint   `gorm:"primarykey"`
	Username          string `gorm:"not null;unique"`
	EncryptedPassword string `gorm:"not null"`
	Videos            []Video
	Comments          []Comment
	Likes             []Video `gorm:"many2many:like"`
	Fans              []User  `gorm:"many2many:follow"`
}

func NewUserDB() *User {
	return &User{}
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Insert(ctx context.Context, username, encpassword string) error {
	result := DB.WithContext(ctx).Create(&User{
		Username: username,
		EncryptedPassword: encpassword,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}