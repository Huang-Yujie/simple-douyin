package user

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func FanList(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.FanListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &userproto.GetFanListReq{
		AppUserId: appUserID,
		UserId:    param.UserId,
	}
	users, err := rpc.GetFanList(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &respond.FanListResp{
		BaseResp: respond.Success,
		UserList: respond.PackUsers(users),
	}
	respond.Send(c, resp)
}
