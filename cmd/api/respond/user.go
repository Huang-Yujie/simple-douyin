package respond

type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
}

type UserRegisterResp struct {
	BaseResp
	UserID int64  `json:"user_id"`
	Token  string `json:"token"` // 用户鉴权token
}

type UserLoginResp struct {
	BaseResp
	UserID int64  `json:"user_id"`
	Token  string `json:"token"` // 用户鉴权token
}

type UserQueryResp struct {
	BaseResp
	User *User `json:"user,omitempty"`
}

type FollowListResp struct {
	BaseResp
	UserList []*User `json:"user_list,omitempty"`
}

type FanListResp struct {
	BaseResp
	UserList []*User `json:"user_list,omitempty"`
}
