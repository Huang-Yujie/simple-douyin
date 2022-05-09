package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"simple-douyin/db_srv/global"
	"simple-douyin/db_srv/model"
	"simple-douyin/db_srv/proto/user"
)
type Server struct {}

func (s *Server) CheckPassword(ctx context.Context, req *proto.CheckPasswordReq) (resp *proto.CheckPasswordResp, err error) {
	return nil, nil
}

func (s *Server) CreateUser(ctx context.Context, req *proto.CreateUserReq) (resp *proto.CreateUserResp, err error) {
	var user model.User
	result := global.DB.Where(&model.User{
		Mobile: req.Mobile,
	}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = "用户" + req.Mobile
	user.Picture = ""
	user.Profile = "此用户还没有填写个人简介"

	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	resp = &proto.CreateUserResp{
		Id: user.ID,
		Password: user.Password,
		Mobile: user.Mobile,
		Nickname: user.NickName,
		Gender: int32(user.Gender),
		Profile: user.Profile,
	}
	return
}

func (s *Server) UpdateUser(ctx context.Context, req *proto.UpdateUserReq) (resp *proto.UpdateUserResp, err error) {
	return nil, nil
}

func (s *Server) GetUserById(ctx context.Context, req *proto.UserIdReq) (resp *proto.UserInfoResq, err error) {
	return nil, nil
}

func (s *Server) GetUserByMobile(ctx context.Context, req *proto.UserMobileReq) (resp *proto.UserInfoResq, err error) {
	return nil, nil
}


