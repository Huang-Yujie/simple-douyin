package service

import (
	"context"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/cmd/user/dal/model"
	"simple-douyin/kitex_gen/userproto"
)

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUser(req *userproto.GetUserReq) (*userproto.UserInfo, error) {
	appUerId := req.AppUserId
	userId := req.UserId

	return s.GetUserInfoByID(appUerId, userId, false)
}

//GetUserInfoByID  查询userId的信息 封装为UserInfo返回，appUerId主要用于判断当前用户是否关注了userId用户
//queryFollow为true 说明查用当前用户的关注列表
func (s *GetUserService) GetUserInfoByID(appUerId, userId int64, queryFollow bool) (*userproto.UserInfo, error) {
	user, err := dal.GetUser(s.ctx, &model.User{UserID: userId})
	if err != nil {
		return nil, err
	}
	u := user[0]
	followCnt, err := dal.GetFollowCount(s.ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	fanCnt, err := dal.GetFanCount(s.ctx, u.UserID)
	if err != nil {
		return nil, err
	}

	//如果不是查当前用户，而且不是查关注列表  就去查询是否关注
	if appUerId != u.UserID && !queryFollow {
		queryFollow, err = dal.IsFollow(s.ctx, appUerId, u.UserID)
		if err != nil {
			return nil, err
		}
	}

	userInfo := &userproto.UserInfo{
		UserId:        u.UserID,
		Username:      u.UserName,
		FollowCount:   followCnt,
		FollowerCount: fanCnt,
		IsFollow:      queryFollow,
	}
	return userInfo, nil
}
