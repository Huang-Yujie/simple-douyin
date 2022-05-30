package service

import (
	"context"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/kitex_gen/userproto"
)

type GetFollowListService struct {
	ctx context.Context
}

func NewGetFollowListService(ctx context.Context) *GetFollowListService  {
	return &GetFollowListService{ctx: ctx}
}

//GetFollowList 两种情况 一种是查看自己的关注列表和粉丝列表  还有一种是查看别人的关注列表和粉丝列表  若appUserId(当前用户id)和userId相等
//说明查看的是当前用户的  若不同 则需要返回当前用户是否关注了点进主页用户
//若是查看当前用户 也就是自己的关注列表  则直接查就行  理论上是全部关注的  is_follow字段表示当前登录用户是否关注了列表中的用户
//若是查看别的用户的关注列表 则需要再查看当前用户是否关注了列表中的用户   粉丝列表类似
func (s *GetFollowListService) GetFollowList(req *userproto.GetFollowListReq) ([]*userproto.UserInfo, error) {
	appUserId := req.AppUserId
	//userId := req.UserId

	//查看当前用户的关注列表
	uids, err := dal.GetFollowList(s.ctx, appUserId)
	if err != nil {
		return nil, err
	}

	userInfos := make([]*userproto.UserInfo, 0)

	for _, uid := range uids {
		uI, err := NewGetUserService(s.ctx).GetUserInfoByID(appUserId, uid, true)
		if err != nil {
			return nil, err
		}

		userInfos = append(userInfos, uI)
	}

	return userInfos, nil
}

