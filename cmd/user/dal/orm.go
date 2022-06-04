package dal

import (
	"context"
	"simple-douyin/cmd/user/dal/model"
)

//GetUserByID 需要通过用户ID查询用户信息
func GetUserByID(ctx context.Context, userID int64) (*model.User, error) {
	var user model.User
	result := DB.WithContext(ctx).Where("user_id = ?", userID).First(&user)
	if result.Error != nil {
		return &user, result.Error
	}
	return &user, nil
}

// GetUserByName 需要通过用户名查询用户信息
func GetUserByName(ctx context.Context, userName string) ([]*model.User, error) {
	var users []*model.User
	result := DB.WithContext(ctx).Where("username = ?", userName).Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}


//CreateUser 根据所给user信息创建用户,并返回用户ID
func CreateUser(ctx context.Context, userName, encPassword string) (int64, error) {
	u := model.User{
		Username: userName,
		EncryptedPassword: encPassword,
	}

	if err := DB.WithContext(ctx).Create(&u).Error; err!=nil {
		return -1, err
	}
	return int64(u.UserID), nil
}

//FollowUser 根据传入的两个用户id  执行用户A关注用户B操作
func FollowUser(ctx context.Context, userID, fanID int64) error {
	follow := model.Follow{
		UserId: uint(userID),
		FanId: uint(fanID),
	}
	if err := DB.WithContext(ctx).Create(&follow).Error; err != nil {
		return err
	}
	return nil
}

//UnFollowUser 根据传入的两个用户id  执行用户A取消关注用户B操作
func UnFollowUser(ctx context.Context, userID, fanID int64) error {
	result := DB.WithContext(ctx).Where("user_user_id = ? AND fan_user_id = ?", userID, fanID).Delete(&model.Follow{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//MGetFanUser 根据userID 获取该用户所关注的用户ID列表
func MGetFanUser(ctx context.Context, fanID int64) ([]int64, error) {
	var follows []*model.Follow
	result := DB.WithContext(ctx).Where("fan_user_id = ?", fanID).Find(&follows)
	if result.Error != nil {
		return []int64{}, result.Error
	}
	userIDs := make([]int64, 0)
	for i:=0; i<len(follows); i++ {
		userIDs = append(userIDs, int64(follows[i].UserId))
	}
	return userIDs, nil
}

//MGetFollowUser 根据userID 获取这个用户的粉丝ID列表
func MGetFollowUser(ctx context.Context, userID int64) ([]int64, error) {
	var follows []*model.Follow
	result := DB.WithContext(ctx).Where("user_user_id = ?", userID).Find(&follows)
	if result.Error != nil {
		return []int64{}, result.Error
	}
	fanIDs := make([]int64, 0)
	for i:=0; i<len(follows); i++ {
		fanIDs = append(fanIDs, int64(follows[i].FanId))
	}
	return fanIDs, nil
}

//GetFanCount 传入用户ID 查询粉丝用户数量
func GetFanCount(ctx context.Context, userID int64) (int64, error)  {
	var count int64
	if err := DB.WithContext(ctx).Model(&model.Follow{}).Where("user_user_id = ?", userID).Count(&count).Error; err!=nil {
		panic(err)
	}
	return count, nil
}

//GetFollowCount 传入用户ID 查询关注用户ID集合
func GetFollowCount(ctx context.Context, fanID int64) (int64, error) {
	var count int64
	if err := DB.WithContext(ctx).Model(&model.Follow{}).Where("fan_user_id = ?", fanID).Count(&count).Error; err!=nil {
		panic(err)
	}
	return count, nil
}

//IsFollow 根据传入两个用户ID 在like表中查询用户A是否关注用户B
func IsFollow(ctx context.Context, userID , fanID int64) (bool, error) {
	var count int64
	err := DB.WithContext(ctx).Model(&model.Follow{}).Where("user_user_id = ? AND fan_user_id = ?", userID, fanID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count>0, nil
}
