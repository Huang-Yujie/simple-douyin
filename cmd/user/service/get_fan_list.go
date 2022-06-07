package service

import (
	"context"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/kitex_gen/userproto"
)

type GetFanListService struct {
	ctx context.Context
}

func NewGetFanListService(ctx context.Context) *GetFanListService {
	return &GetFanListService{ctx: ctx}
}

func (s *GetFanListService) GetFanList(req *userproto.GetFanListReq) ([]*userproto.UserInfo, error) {
	appUserId := req.AppUserId
	userId := req.UserId

	//查看当前用户的粉丝列表uids
	uids, err := dal.MGetFanUser(s.ctx, userId)
	if err != nil {
		return nil, err
	}
	if len(uids) == 0 {
		return nil, nil
	}
	userInfos := make([]*userproto.UserInfo, len(uids))

	for i, uid := range uids {
		userInfo, err := NewGetUserService(s.ctx).GetUserInfoByID(appUserId, uid)
		if err != nil {
			return nil, err
		}
		userInfos[i] = userInfo
	}

	return userInfos, nil
}
