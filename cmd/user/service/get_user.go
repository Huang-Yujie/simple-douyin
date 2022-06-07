package service

import (
	"context"
	"errors"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/errno"

	"gorm.io/gorm"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *userproto.GetUserReq) (*userproto.UserInfo, error) {
	return s.GetUserInfoByID(req.AppUserId, req.UserId)
}

//GetUserInfoByID  查询userId的信息 封装为UserInfo返回，appUerId主要用于判断当前用户是否关注了userId用户
func (s *GetUserService) GetUserInfoByID(appUserId, userId int64) (*userproto.UserInfo, error) {
	user, err := dal.GetUserByID(s.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return nil, errno.UserNotExistErr
		}
		return nil, err
	}

	followCnt, err := dal.GetFollowCount(s.ctx, int64(user.UserID))
	if err != nil {
		return nil, err
	}

	fanCnt, err := dal.GetFanCount(s.ctx, int64(user.UserID))
	if err != nil {
		return nil, err
	}

	isFollow := false  // 默认为false
	if appUserId > 0 { // 如果已登录则查询
		isFollow, err = dal.IsFollow(s.ctx, appUserId, int64(user.UserID))
		if err != nil {
			return nil, err
		}
	}

	userInfo := &userproto.UserInfo{
		UserId:        int64(user.UserID),
		Username:      user.Username,
		FollowCount:   followCnt,
		FollowerCount: fanCnt,
		IsFollow:      isFollow,
	}
	return userInfo, nil
}
