package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"
	"simple-douyin/cmd/pkg/errno"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/cmd/user/model"
	"simple-douyin/kitex_gen/userproto"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *userproto.CheckUserReq) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.UserAccount.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.UserAccount.Username
	users, err := dal.QueryUser(s.ctx, &model.User{UserName: userName})
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.EncryptedPassword != passWord {
		return 0, errno.LoginErr
	}
	return u.UserID, nil
}
