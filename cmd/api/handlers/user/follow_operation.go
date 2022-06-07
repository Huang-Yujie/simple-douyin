package user

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func FollowOperation(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.FollowOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // 关注
		req := &userproto.FollowUserReq{
			FanUserId:      appUserID,
			FollowedUserId: param.ToUserId,
		}
		if err := rpc.FollowUser(c, req); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	case 2: // 取关
		req := &userproto.UnFollowUserReq{
			FanUserId:      appUserID,
			FollowedUserId: param.ToUserId,
		}
		if err := rpc.UnFollowUser(c, req); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errno.ParamErr)
	}
}
