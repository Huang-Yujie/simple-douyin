package main

import (
	"context"
	userproto2 "simple-douyin/kitex_gen/userproto"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userproto2.CreateUserReq) (resp *userproto2.CreateUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userproto2.GetUserReq) (resp *userproto2.GetUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userproto2.CheckUserReq) (resp *userproto2.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *userproto2.FollowUserReq) (resp *userproto2.FollowUserResp, err error) {
	// TODO: Your code here...
	return
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *userproto2.UnFollowUserReq) (resp *userproto2.UnFollowUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userproto2.GetFollowListReq) (resp *userproto2.GetFollowListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFanList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFanList(ctx context.Context, req *userproto2.GetFanListReq) (resp *userproto2.GetFanListResp, err error) {
	// TODO: Your code here...
	return
}
