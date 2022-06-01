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

func CommentList(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.CommentListReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	req := &videoproto.GetCommentsReq{
		AppUserId: appUserID,
		VideoId:   param.VideoId,
	}
	comments, err := rpc.GetComments(c, req)
	if err != nil {
		respond.Error(c, err)
		return
	}
	n := len(comments)
	authors := make([]*userproto.UserInfo, n)
	for i := 0; i < n; i++ {
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    comments[i].UserId,
		}
		authors[i], err = rpc.GetUser(c, subReq)
		if err != nil {
			respond.Error(c, err)
			return
		}
	}
	resp := &respond.CommentListResp{
		BaseResp:    respond.Success,
		CommentList: respond.PackComments(comments, authors),
	}
	respond.Send(c, resp)
}
