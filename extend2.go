package forms

//关注操作
type FollowOperation struct {
	User_Id     string `json:"user_id" binding:"required,max=32"`
	Token       string `json:"token" binding:"required"`
	To_User_Id  string `json:"to_user_id" binding:"required,max=32"`
	Action_Type string `json:"action_type" binding:"required" msg:"1-关注，2-取消关注"`
}

//关注列表
type FollowList struct {
	User_Id string `json:"user_id" binding:"required,max=32"`
	Token   string `json:"token" binding:"required"`
}

//粉丝列表
type FanswList struct {
	User_Id string `json:"user_id" binding:"required,max=32"`
	Token   string `json:"token" binding:"required"`
}
