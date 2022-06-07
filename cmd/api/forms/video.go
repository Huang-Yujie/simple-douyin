package forms

//视频流
type VideoFeedReq struct {
	LatestTime int64 `form:"latest_time" json:"latest_time"`
}

//点赞操作
type LikeOperationReq struct {
	VideoId    int64 `form:"video_id" json:"video_id" binding:"required"`
	ActionType int   `form:"action_type" json:"action_type" binding:"required" msg:"1-点赞，2-取消点赞"`
}

//点赞列表
type LikeListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

//评论操作
type CommentOperationReq struct {
	VideoId     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int    `form:"action_type" json:"action_type" binding:"required" msg:"1-发布评论，2-删除评论"`
	CommentText string `form:"comment_text" json:"comment_text" msg:"action_type==1时使用"`
	CommentId   int64  `form:"comment_id" json:"comment_id" msg:"要删除的评论id，action_type==2时使用"`
}

//评论列表
type CommentListReq struct {
	VideoId int64 `form:"video_id" json:"video_id" binding:"required"`
}

//投稿接口
type VideoUploadReq struct {
	Title string `form:"title" json:"title" binding:"required"`
}

//发布列表
type VideoListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}
