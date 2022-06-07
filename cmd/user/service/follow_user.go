package service

import (
	"context"
	"simple-douyin/cmd/user/dal"
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
	return dal.FollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}

//UnFollowUser unFollow user by id
func (s *FollowUserService) UnFollowUser(req *userproto.UnFollowUserReq) error {
	return dal.UnFollowUser(s.ctx, int64(req.FanUserId), int64(req.FollowedUserId))
}
