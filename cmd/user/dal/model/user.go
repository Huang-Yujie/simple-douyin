package model

type User struct {
	UserID            uint   `gorm:"primarykey"`
	Username          string `gorm:"not null;unique"`
	EncryptedPassword string `gorm:"not null"`
	Videos            []*Video
	Comments          []*Comment
	Likes             []*Video `gorm:"many2many:like;joinForeignKey:user_id;joinReferences:video_id;"`
	Fans              []*User  `gorm:"many2many:follow;joinForeignKey:user_id;joinReferences:fan_id;"`
}

func (u *User) TableName() string {
	return "user"
}
