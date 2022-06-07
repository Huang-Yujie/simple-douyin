package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/pkg/errno"

	"simple-douyin/kitex_gen/userproto"

	"gorm.io/gorm"
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
	_, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err == nil { // 找到了
			return 0, errno.UserAlreadyExistErr
		}
		return 0, err // 其他错误
	}
	h := md5.New()
	if _, err := io.WriteString(h, req.UserAccount.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return dal.CreateUser(s.ctx, req.UserAccount.Username, passWord)
}
