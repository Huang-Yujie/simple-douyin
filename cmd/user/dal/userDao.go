package dal

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"simple-douyin/cmd/user/model"
)

//QueryUser 需要通过 用户ID 或者 用户名 查询用户信息
func QueryUser(ctx context.Context, user *model.User) ([]*model.User, error) {
	var targetUser []*model.User
	err:=DB.WithContext(ctx).Where(&user).Find(&targetUser).Error
	return targetUser, err
}

//CreateUser create user info  根据所给user信息创建用户
func CreateUser(ctx context.Context, user *model.User) (int64, error) {
	err := DB.WithContext(ctx).Create(&user).Error
	return user.UserID, err
}

//GetFollowCount 根据userID获取关注总数
func GetFollowCount(ctx context.Context, userID int64) (int64, error) {
	//建议使用sql来查询
	var count int64
	rows, err := DB.WithContext(ctx).Table("follow").Where("user_user_id = ?", userID).Select("fan_user_id").Rows()
	defer rows.Close()
	for rows.Next() {
		count++
	}
	return count, err
}

//GetFollowerCount 根据userID获取粉丝总数
func GetFollowerCount(ctx context.Context, userID int64) (int64, error) {
	//建议使用sql来查询
	var count int64
	rows, err := DB.WithContext(ctx).Table("follow").Where("fan_user_id = ?", userID).Select("user_user_id").Rows()
	defer rows.Close()
	for rows.Next() {
		count++
	}
	return count, err
}

//GetFollowList  根据userID 获取该用户  所关注的用户ID列表
func GetFollowList(ctx context.Context, userID int64) ([]int, error) {
	favs:=[]int{}
	rows, err := DB.WithContext(ctx).Table("follow").Where("user_user_id = ?", userID).Select("fan_user_id").Rows()
	defer rows.Close()
	for rows.Next() {
		var id int
		rows.Scan(&id)
		favs = append(favs, id)
	}
	return favs, err
}

//GetFollowerList  根据userID 获取这个用户的粉丝ID列表
func GetFollowerList(ctx context.Context, userID int64) ([]int, error) {
	fans:=[]int{}
	rows, err := DB.WithContext(ctx).Table("follow").Where("fan_user_id = ?", userID).Select("user_user_id").Rows()
	defer rows.Close()
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fans = append(fans, id)
	}
	return fans, err
}

//IsFollow  根据传入的两个用户id  判断userA是否关注userB
func IsFollow(ctx context.Context, userA , userB int64) (bool, error) {
	rows, err := DB.WithContext(ctx).Table("follow").Where("user_user_id = ? AND fan_user_id= ?", userA, userB).Rows()
	flag := !errors.Is(err, gorm.ErrRecordNotFound)
	defer rows.Close()
	return flag, nil
}

//Follow 根据传入的两个用户id  执行用户A关注用户B操作
func Follow(ctx context.Context, userA, userB *model.User) error {
	DB.Model(&userB).WithContext(ctx).Association("Fans").Append(userA)
	return nil
}

//UnFollow 根据传入的两个用户id  执行用户A取消关注用户B操作
func UnFollow(ctx context.Context, userA, userB *model.User) error {
	DB.Model(&userB).WithContext(ctx).Association("Fans").Delete(userA)
	return nil
}

