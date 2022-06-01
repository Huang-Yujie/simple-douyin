package video

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func LikeOperation(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.LikeOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // 点赞
		req := &videoproto.LikeVideoReq{
			UserId:  appUserID,
			VideoId: param.VideoId,
		}
		if err := rpc.LikeVideo(c, req); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	case 2: // 取消点赞
		req := &videoproto.UnLikeVideoReq{
			UserId:  appUserID,
			VideoId: param.VideoId,
		}
		if err := rpc.UnLikeVideo(c, req); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errno.ParamErr)
	}
}
