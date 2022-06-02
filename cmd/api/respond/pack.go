package respond

import (
	"simple-douyin/cmd/api/cache"
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

func PackVideo(v *videoproto.VideoInfo, author *userproto.UserInfo) (*Video, error) {
	playURL, err := cache.GetPlayURL(v.VideoBaseInfo.PlayAddr)
	if err != nil {
		return nil, err
	}
	return &Video{
		ID:           v.VideoId,
		Author:       PackUser(author),
		PlayAddr:     v.VideoBaseInfo.PlayAddr,
		CoverAddr:    playURL,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		IsFavorite:   v.IsFavorite,
		Title:        v.VideoBaseInfo.Title,
	}, nil
}

func PackVideos(vs []*videoproto.VideoInfo, authors []*userproto.UserInfo) ([]*Video, error) {
	n := len(vs)
	videos := make([]*Video, n)
	for i := 0; i < n; i++ {
		var err error
		videos[i], err = PackVideo(vs[i], authors[i])
		if err != nil {
			return nil, err
		}
	}
	return videos, nil
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
