package user

import (
	"simple-douyin/cmd/api/auth"
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	param := new(forms.UserRegisterReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &userproto.CreateUserReq{
		UserAccount: &userproto.UserAccount{
			Username: param.Username,
			Password: param.Password,
		},
	}
	userID, err := rpc.CreateUser(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	token, err := auth.GenerateToken(userID)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &respond.UserRegisterResp{
		BaseResp: respond.Success,
		UserID:   userID,
		Token:    token,
	}
	respond.Send(c, resp)
}
