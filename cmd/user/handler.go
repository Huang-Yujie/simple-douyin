package main

import (
	"context"
	"simple-douyin/cmd/user/pack"
	"simple-douyin/cmd/user/service"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userproto.CreateUserReq) (resp *userproto.CreateUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.CreateUserResp)

	if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userID, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = userID
	return resp, nil
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *userproto.GetUserReq) (resp *userproto.GetUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.GetUserResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	userInfo, err := service.NewGetUserService(ctx).GetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserInfo = userInfo
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userproto.CheckUserReq) (resp *userproto.CheckUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.CheckUserResp)

	if len(req.UserAccount.Username) == 0 || len(req.UserAccount.Password) == 0 || len(req.UserAccount.Username) > 32 || len(req.UserAccount.Password) > 32 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) FollowUser(ctx context.Context, req *userproto.FollowUserReq) (resp *userproto.FollowUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.FollowUserResp)

	if req.FollowedUserId < 0 || req.FanUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFollowUserService(ctx).FollowUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// UnFollowUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) UnFollowUser(ctx context.Context, req *userproto.UnFollowUserReq) (resp *userproto.UnFollowUserResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.UnFollowUserResp)

	if req.FollowedUserId < 0 || req.FanUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewFollowUserService(ctx).UnFollowUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFollowList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollowList(ctx context.Context, req *userproto.GetFollowListReq) (resp *userproto.GetFollowListResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.GetFollowListResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	follows, err := service.NewGetFollowListService(ctx).GetFollowList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserInfos = follows
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFanList implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFanList(ctx context.Context, req *userproto.GetFanListReq) (resp *userproto.GetFanListResp, err error) {
	// TODO: Your code here...
	resp = new(userproto.GetFanListResp)

	if req.UserId < 0 || req.AppUserId < 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	fans, err := service.NewGetFanListService(ctx).GetFanList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.UserInfos = fans
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
