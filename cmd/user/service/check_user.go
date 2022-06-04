package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"simple-douyin/cmd/user/dal"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/errno"

	"gorm.io/gorm"
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
	user, err := dal.GetUserByName(s.ctx, req.UserAccount.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果没找到
			return 0, errno.UserNotExistErr
		}
		return 0, err
	}
	if user.EncryptedPassword != passWord {
		return 0, errno.LoginErr
	}
	return int64(user.UserID), nil
}
