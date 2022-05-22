package forms

//点赞操作
type LikeOperation struct {
	User_Id     string `json:"user_id" binding:"required,max=32"`
	Token       string `json:"token" binding:"required"`
	Video_Id    string `json:"video_id" binding:"required"`
	Action_Type string `json:"action_type" binding:"required" msg:"1-点赞，2-取消点赞"`
}

//点赞列表
type LikeList struct {
	User_Id string `json:"user_id" binding:"required,max=32"`
	Token   string `json:"token" binding:"required"`
}

//评论操作
type CommentOperation struct {
	User_Id     string `json:"user_id" binding:"required,max=32"`
	Token       string `json:"token" binding:"required"`
	Video_Id    string `json:"video_id" binding:"required"`
	Action_Type string `json:"action_type" binding:"required" msg:"1-发布评论，2-删除评论"`
	Commen_Text string `json:"commen_text" msg:"action_type==1时使用"`
	Commen_Id   string `json:"commen_id" msg:"要删除的评论id，action_type==2时使用"`
}

//评论列表
type CommentList struct {
	Token    string `json:"token" binding:"required"`
	Video_Id string `json:"video_id" binding:"required"`
}
