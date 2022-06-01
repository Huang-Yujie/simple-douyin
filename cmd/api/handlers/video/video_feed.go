package video

import (
	"simple-douyin/cmd/api/auth"
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"
	"time"

	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	appUserID, err := auth.GetUserID(c)
	if err != nil { // 用户未登录
		appUserID = -1
	}
	param := new(forms.VideoFeedReq)
	if err := c.ShouldBind(param); err != nil {
		param.LatestTime = time.Now().Unix()
	}
	req := &videoproto.GetVideosByTimeReq{
		AppUserId:  appUserID,
		LatestTime: param.LatestTime,
		Count:      constants.FeedCount,
	}
	videos, nextTime, err := rpc.GetVideosByTime(c, req)
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
	resp := &respond.VideoFeedResp{
		BaseResp:  respond.Success,
		NextTime:  nextTime,
		VideoList: respond.PackVideos(videos, authors),
	}
	respond.Send(c, resp)

}
