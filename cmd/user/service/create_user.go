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

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *userproto.CreateUserReq) (int64, error) {
	user, err := dal.QueryUser(s.ctx, &model.User{UserName: req.UserAccount.Username})
	if err != nil {
		return -1, err
	}
	if len(user) != 0 {
		return -1, errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err := io.WriteString(h, req.UserAccount.Password); err != nil {
		return -1, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	moudelUser := &model.User{UserName: req.UserAccount.Username, EncryptedPassword: passWord}
	return dal.CreateUser(s.ctx, moudelUser)

}
