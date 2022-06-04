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
	for i, u := range us {
		users[i] = PackUser(u)
	}
	return users
}

func PackVideo(v *videoproto.VideoInfo, author *userproto.UserInfo) (*Video, error) {
	playURL, err := cache.GetPlayURL(v.VideoBaseInfo.OssVideoId)
	if err != nil {
		return nil, err
	}
	// coverURL, err := cache.GetCoverURL(v.VideoBaseInfo.OssVideoId)
	// if err != nil {
	// 	return nil, err
	// }
	coverURL := "https://tva1.sinaimg.cn/large/e6c9d24ely1h2wrrikp8uj20tc1io422.jpg"
	return &Video{
		ID:           v.VideoId,
		Author:       PackUser(author),
		PlayAddr:     playURL,
		CoverAddr:    coverURL,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		IsFavorite:   v.IsFavorite,
		Title:        v.VideoBaseInfo.Title,
	}, nil
}

func PackVideos(vs []*videoproto.VideoInfo, authors []*userproto.UserInfo) ([]*Video, error) {
	videos := make([]*Video, len(vs))
	for i, v := range vs {
		var err error
		videos[i], err = PackVideo(v, authors[i])
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

func PackComments(cs []*videoproto.CommentInfo, authors []*userproto.UserInfo) []*Comment {
	comments := make([]*Comment, len(cs))
	for i, c := range cs {
		comments[i] = PackComment(c, authors[i])
	}
	return comments
}
