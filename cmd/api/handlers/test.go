// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package handlers

import (
	"context"
	"strconv"

	"simple-douyin/pkg/errno"

	"simple-douyin/kitex_gen/videoproto"

	"simple-douyin/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// CreateNote create note info
func CreateVideo(c *gin.Context) {
	var videoVar VideoParam
	if err := c.ShouldBind(&videoVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	// claims := jwt.ExtractClaims(c)
	// userID := int64(claims[constants.IdentityKey].(float64))
	videoBaseInfo := &videoproto.VideoBaseInfo{
		UserId:    videoVar.UserId,
		PlayAddr:  videoVar.PlayAddr,
		CoverAddr: videoVar.CoverAddr,
		Title:     videoVar.Title,
	}

	err := rpc.CreateVideo(context.Background(), &videoproto.CreateVideoReq{
		VideoBaseInfo: videoBaseInfo,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

func QueryVideoByUid(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	appUserId, err := strconv.ParseInt(c.Query("appUserId"), 10, 64)

	getVideosByUserIdReq := &videoproto.GetVideosByUserIdReq{
		AppUserId: appUserId,
		UserId:    uid,
	}
	resp, err := rpc.QueryVideoByUid(context.Background(), getVideosByUserIdReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.GET("/queryByTime", handlers.QueryVideoByTime)
func QueryVideoByTime(c *gin.Context) {
	appUserId, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	latestTime, err := strconv.ParseInt(c.Query("latestTime"), 10, 64)
	count, err := strconv.ParseInt(c.Query("count"), 10, 64)
	getVideosByTimeReq := &videoproto.GetVideosByTimeReq{
		AppUserId:  appUserId,
		LatestTime: latestTime,
		Count:      count,
	}
	resp, err := rpc.QueryVideoByTime(context.Background(), getVideosByTimeReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.POST("/like", handlers.LikeVideo)
func LikeVideo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	vid, err := strconv.ParseInt(c.Query("vid"), 10, 64)
	likeVideoReq := &videoproto.LikeVideoReq{
		UserId:  userId,
		VideoId: vid,
	}
	resp, err := rpc.LikeVideo(context.Background(), likeVideoReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.DELETE("/unlike", handlers.UnlikeVideo)
func UnlikeVideo(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	vid, err := strconv.ParseInt(c.Query("vid"), 10, 64)
	unLikeVideoReq := &videoproto.UnLikeVideoReq{
		UserId:  userId,
		VideoId: vid,
	}
	resp, err := rpc.UnLikeVideo(context.Background(), unLikeVideoReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.POST("/creatcomment", handlers.CreatComment)
func CreateComment(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	vid, err := strconv.ParseInt(c.Query("vid"), 10, 64)
	content := c.Query("content")
	createCommentReq := &videoproto.CreateCommentReq{
		UserId:  userId,
		VideoId: vid,
		Content: content,
	}
	resp, err := rpc.CreateComment(context.Background(), createCommentReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.DELETE("deletecomment", handlers.DeleteComment)
func DeleteComment(c *gin.Context) {
	commentId, err := strconv.ParseInt(c.Query("commentId"), 10, 64)
	deleteCommentReq := &videoproto.DeleteCommentReq{
		CommentId: commentId,
	}
	resp, err := rpc.DeleteComment(context.Background(), deleteCommentReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}

// video1.GET("/getcomment", handlers.GetComments)
func GetComments(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("uid"), 10, 64)
	vid, err := strconv.ParseInt(c.Query("vid"), 10, 64)
	getCommentsReq := &videoproto.GetCommentsReq{
		AppUserId: userId,
		VideoId:   vid,
	}
	resp, err := rpc.GetComments(context.Background(), getCommentsReq)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
