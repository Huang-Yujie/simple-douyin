package forms

import "google.golang.org/protobuf/compiler/protogen"

//视频流
type VideoStream struct {
	LatestTime string `form:"latesttime" json:"latesttime"`
	Token      string `form:"token" json:"token"`
}

//点赞操作
type LikeOperation struct {
	UserId     string `form:"userid" json:"userid" binding:"required,max=32"`
	Token      string `form:"token" json:"token" binding:"required"`
	VideoId    string `form:"videoid" json:"videoid" binding:"required"`
	ActionType string `form:"actiontype" json:"actiontype" binding:"required" msg:"1-点赞，2-取消点赞"`
}

//点赞列表
type LikeList struct {
	UserId string `form:"userid" json:"userid" binding:"required,max=32"`
	Token  string `form:"token" json:"token" binding:"required"`
}

//评论操作
type CommentOperation struct {
	UserId      string `form:"userid" json:"userid" binding:"required,max=32"`
	Token       string `form:"token" json:"token" binding:"required"`
	VideoId     string `form:"videoid" json:"videoid" binding:"required"`
	ActionType  string `form:"actiontype" json:"actiontype" binding:"required" msg:"1-发布评论，2-删除评论"`
	CommentText string `form:"commentext" json:"commentext" msg:"action_type==1时使用"`
	CommentId   string `form:"commentid" json:"commenid" msg:"要删除的评论id，action_type==2时使用"`
}

//评论列表
type CommentList struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoId string `form:"videoid" json:"videoid" binding:"required"`
}

//投稿接口
type UploadInterface struct {
	Data  protogen.File `form:"data" json:"data" binding:"required"`
	Token string        `form:"token" json:"token" binding:"required"`
	Title string        `form:"title" json:"title" binding:"required"`
}

//发布列表
type ReleaseList struct {
	Token  string `form:"token" json:"token" binding:"required"`
	UserId string `form:"userid" json:"userid" binding:"required,max=32"`
}
