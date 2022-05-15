package db

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `gorm:"column:email;type:varchar(100);not null;unique"`
	Password   string `gorm:"column:password;type:varchar(200);not null;unique"`
	Nickname   string `gorm:"column:password;type:varchar(20);"`
	Picture    string `gorm:"column:picture;type:varchar(200);"`
	Profile    string `gorm:"column:profile;type:varchar(200)"`
}

func(u *User) TableName() string {
	return "user"
}

func(u *User) CreateUser(ctx context.Context, email, password string) error {
	user := &User{
		Email: email,
		Password: password,
	}
	return DB.WithContext(ctx).Create(user).Error
}

func(u *User) GetUserById(ctx context.Context, id uint) (*User, error) {
	user := &User{}
	if err := DB.WithContext(ctx).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func(u *User) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	if err := DB.WithContext(ctx).Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func(u *User) UpdateUser(ctx context.Context, user User) error {
	DB.WithContext(ctx).Updates(user)
	return nil
}



