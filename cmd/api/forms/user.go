package forms

//用户注册
type UserRegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32" msg:"最长32个字符串"`
	Password string `form:"password" json:"password" binding:"required,max=32" msg:"最长32个字符串"`
}

//用户登录
type UserLoginReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

//用户信息
type UserQueryReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

//关注操作
type FollowOperationReq struct {
	ToUserId   int64 `form:"to_user_id" json:"to_user_id" binding:"required,max=32"`
	ActionType int   `form:"action_type" json:"action_type" binding:"required" msg:"1-关注，2-取消关注"`
}

//关注列表
type FollowListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}

//粉丝列表
type FanListReq struct {
	UserId int64 `form:"user_id" json:"user_id" binding:"required,max=32"`
}
