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

// UpdateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *userproto2.UpdateUserReq) (resp *userproto2.UpdateUserResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, req *userproto2.GetUserByIdReq) (resp *userproto2.GetUserByIdResp, err error) {
	// TODO: Your code here...
	return
}

// GetUserByEmail implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, req *userproto2.GetUserByEmailReq) (resp *userproto2.GetUserByEmailResp, err error) {
	// TODO: Your code here...
	return
}

// CheckPassword implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckPassword(ctx context.Context, req *userproto2.CheckPasswordReq) (resp *userproto2.CheckPasswordResp, err error) {
	// TODO: Your code here...
	return
}
