package video

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/errno"

	"github.com/gin-gonic/gin"
)

func CommentOperation(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.CommentOperationReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	switch param.ActionType {
	case 1: // 评论
		req := &videoproto.CreateCommentReq{
			UserId:  appUserID,
			VideoId: param.VideoId,
			Content: param.CommentText,
		}
		comment, err := rpc.CreateComment(c, req)
		if err != nil {
			respond.Error(c, err)
			return
		}
		subReq := &userproto.GetUserReq{
			AppUserId: appUserID,
			UserId:    comment.UserId,
		}
		author, err := rpc.GetUser(c, subReq)
		if err != nil {
			respond.Error(c, err)
			return
		}
		resp := &respond.CreateCommentResp{
			BaseResp: respond.Success,
			Comment:  respond.PackComment(comment, author),
		}
		respond.Send(c, resp)
	case 2: // 删除评论
		req := &videoproto.DeleteCommentReq{
			CommentId: param.CommentId,
		}
		if err := rpc.DeleteComment(c, req); err != nil {
			respond.Error(c, err)
			return
		}
		respond.OK(c)
	default:
		respond.Error(c, errno.ParamErr)
	}
}
