package main

import (
	"context"
	"simple-douyin/kitex_gen/userproto"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userproto.CreateUserReq) (resp *userproto.CreateUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userproto.GetUserReq) (resp *userproto.GetUserResp, err error) {
	// TODO: Your code here...
	return
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userproto.CheckUserReq) (resp *userproto.CheckUserResp, err error) {
	// TODO: Your code here...
	return
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *userproto.FollowUserReq) (resp *userproto.FollowUserResp, err error) {
	// TODO: Your code here...
	return
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *userproto.UnFollowUserReq) (resp *userproto.UnFollowUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserRelations implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserRelations(ctx context.Context, req *userproto.GetUserRelationsReq) (resp *userproto.GetUserRelationsResp, err error) {
	// TODO: Your code here...
	return
}

// GetWhetherBeFollowed implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetWhetherBeFollowed(ctx context.Context, req *userproto.GetWhetherBeFollowedReq) (resp *userproto.GetWhetherBeFollowedResp, err error) {
	// TODO: Your code here...
	return
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userproto.GetFollowListReq) (resp *userproto.GetFollowListResp, err error) {
	// TODO: Your code here...
	return
}

// GetFanList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFanList(ctx context.Context, req *userproto.GetFanListReq) (resp *userproto.GetFanListResp, err error) {
	// TODO: Your code here...
	return
}
