package forms

import "google.golang.org/protobuf/compiler/protogen"

//视频流
type VideoStream struct {
	Latest_Time string `json:"latest_time"`
	Token       string `json:"token"`
}

//用户注册
type UserRegister struct {
	Username string `json:"username" binding:"required,max=32" mas:"最长32个字符串"`
	Password string `json:"password" binding:"required,max=32" mas:"最长32个字符串"`
}

//用户登录
type UserLogin struct {
	Username string `json:"username" binding:"required,max=32"`
	Password string `json:"password" binding:"required,max=32"`
}

//用户信息
type UserInfo struct {
	User_Id string `json:"user_id" binding:"required,max=32"`
	Token   string `json:"token" binding:"required"`
}

//投稿接口
type UploadInterface struct {
	Data  protogen.File `json:"data" binding:"required"`
	Token string        `json:"token" binding:"required"`
	Title string        `json:"title" binding:"required"`
}

//发布列表
type ReleaseList struct {
	Token   string `json:"token" binding:"required"`
	User_Id string `json:"user_id" binding:"required,max=32"`
}
