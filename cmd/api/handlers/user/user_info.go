package user

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.UserQueryReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &userproto.GetUserReq{
		AppUserId: appUserID,
		UserId:    param.UserId,
	}
	user, err := rpc.GetUser(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &respond.UserQueryResp{
		BaseResp: respond.Success,
		User:     respond.PackUser(user),
	}
	respond.Send(c, resp)
}
