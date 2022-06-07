package user

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func FollowList(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.FollowListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &userproto.GetFollowListReq{
		AppUserId: appUserID,
		UserId:    param.UserId,
	}
	users, err := rpc.GetFollowList(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	resp := &respond.FollowListResp{
		BaseResp: respond.Success,
		UserList: respond.PackUsers(users),
	}
	respond.Send(c, resp)
}
