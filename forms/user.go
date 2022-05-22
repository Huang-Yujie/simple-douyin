package forms

//用户注册
type UserRegister struct {
	Username string `form:"username" json:"username" binding:"required,max=32" msg:"最长32个字符串"`
	Password string `form:"password" json:"password" binding:"required,max=32" msg:"最长32个字符串"`
}

//用户登录
type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

//用户信息
type UserInfo struct {
	UserId string `form:"userid" json:"userid" binding:"required,max=32"`
	Token  string `form:"token" json:"token" binding:"required"`
}

//关注操作
type FollowOperation struct {
	UserId     string `form:"userid" json:"userid" binding:"required,max=32"`
	Token      string `form:"token" json:"token" binding:"required"`
	ToUserId   string `form:"touserid" json:"touserid" binding:"required,max=32"`
	ActionType string `form:"actiontype" json:"actiontype" binding:"required" msg:"1-关注，2-取消关注"`
}

//关注列表
type FollowList struct {
	UserId string `form:"userid" json:"userid" binding:"required,max=32"`
	Token  string `form:"token" json:"token" binding:"required"`
}

//粉丝列表
type FansList struct {
	UserId string `form:"userid" json:"userid" binding:"required,max=32"`
	Token  string `form:"token" json:"token" binding:"required"`
}
