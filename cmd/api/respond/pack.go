package respond

import (
	"simple-douyin/kitex_gen/userproto"
	"simple-douyin/kitex_gen/videoproto"
)

func PackUser(u *userproto.UserInfo) *User {
	return &User{
		ID:            u.UserId,
		Name:          u.Username,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func PackUsers(us []*userproto.UserInfo) []*User {
	n := len(us)
	users := make([]*User, n)
	for i := 0; i < n; i++ {
		users[i] = PackUser(us[i])
	}
	return users
}

func PackVideo(v *videoproto.VideoInfo, author *userproto.UserInfo) *Video {
	return &Video{
		ID:           v.VideoId,
		Author:       PackUser(author),
		PlayAddr:     v.VideoBaseInfo.PlayAddr,
		CoverAddr:    v.VideoBaseInfo.CoverAddr,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		IsFavorite:   v.IsFavorite,
		Title:        v.VideoBaseInfo.Title,
	}
}

func PackVideos(vs []*videoproto.VideoInfo, authors []*userproto.UserInfo) []*Video {
	n := len(vs)
	videos := make([]*Video, n)
	for i := 0; i < n; i++ {
		videos[i] = PackVideo(vs[i], authors[i])
	}
	return videos
}

func PackComment(c *videoproto.CommentInfo, author *userproto.UserInfo) *Comment {
	return &Comment{
		ID:         c.CommentId,
		User:       PackUser(author),
		Content:    c.Content,
		CreateDate: c.CreateDate,
	}
}

func PackComments(vs []*videoproto.CommentInfo, authors []*userproto.UserInfo) []*Comment {
	n := len(vs)
	comments := make([]*Comment, n)
	for i := 0; i < n; i++ {
		comments[i] = PackComment(vs[i], authors[i])
	}
	return comments
}
