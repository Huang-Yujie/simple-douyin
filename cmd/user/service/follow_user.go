package service

import (
	"context"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/cmd/user/model"
	"simple-douyin/kitex_gen/userproto"
)

type FollowUserService struct {
	ctx context.Context
}

// NewFollowUserService new CheckUserService
func NewFollowUserService(ctx context.Context) *FollowUserService {
	return &FollowUserService{
		ctx: ctx,
	}
}

//FollowUser Follow user by id
func (s *FollowUserService) FollowUser(req *userproto.FollowUserReq) error {
	userA := req.FanUserId
	userB := req.FollowedUserId
	return dal.Follow(s.ctx, &model.User{UserID: userA}, &model.User{UserID: userB})
}

//UnFollowUser unFollow user by id
func (s *FollowUserService) UnFollowUser(req *userproto.UnFollowUserReq) error {
	userA := req.FanUserId
	userB := req.FollowedUserId
	return dal.UnFollow(s.ctx, &model.User{UserID: userA}, &model.User{UserID: userB})
}
