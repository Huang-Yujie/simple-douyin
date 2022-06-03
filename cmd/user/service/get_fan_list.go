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
	//userId := req.UserId

	//查看当前用户的粉丝列表ids
	uids, err := dal.GetFanList(s.ctx, appUserId)
	if err != nil {
		return nil, err
	}
	userInfos := make([]*userproto.UserInfo, 0)

	for _, uid := range uids {
		uI, err := NewGetUserService(s.ctx).GetUserInfoByID(appUserId, uid, false)
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, uI)
	}

	return userInfos, nil
}
