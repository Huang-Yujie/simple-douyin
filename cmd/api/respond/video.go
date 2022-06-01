package respond

type Video struct {
	ID           int64  `json:"id"`
	Author       *User  `json:"author"`
	PlayAddr     string `json:"play_url"`
	CoverAddr    string `json:"cover_url"`
	LikeCount    int64  `json:"favorite_count"`
	CommentCount int64  `json:"comment_count"`
	IsFavorite   bool   `json:"is_favorite"`
	Title        string `json:"title"`
}

type Comment struct {
	ID         int64  `json:"id"`          // 评论id
	User       *User  `json:"user"`        // 评论用户信息
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type VideoFeedResp struct {
	BaseResp
	NextTime  int64    `json:"next_time"`
	VideoList []*Video `json:"video_list"`
}

type LikeListResp struct {
	BaseResp
	VideoList []*Video `json:"video_list"`
}

type CreateCommentResp struct {
	BaseResp
	Comment *Comment `json:"comment"`
}

type CommentListResp struct {
	BaseResp
	CommentList []*Comment `json:"comment"`
}

type VideoListResp struct {
	BaseResp
	VideoList []*Video `json:"video_list"`
}
