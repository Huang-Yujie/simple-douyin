package video

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.VideoListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &videoproto.GetVideosByUserIdReq{
		AppUserId: appUserID,
		UserId:    param.UserId,
	}
	videos, err := rpc.GetVideosByUserId(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	n := len(videos)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    videos[i].VideoBaseInfo.UserId,
		}
		authors[i], err = rpc.GetUser(c, subReq)
		if err != nil {
			respond.Error(c, err)
			return
		}
	}
	packedVideos, err := respond.PackVideos(videos, authors)
	if err != nil {
		respond.Error(c, err)
	}
	resp := &respond.VideoListResp{
		BaseResp:  respond.Success,
		VideoList: packedVideos,
	}
	respond.Send(c, resp)
}
