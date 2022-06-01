package handlers

import (
	"simple-douyin/cmd/api/auth"
	"simple-douyin/cmd/api/handlers/user"
	"simple-douyin/cmd/api/handlers/video"
	"simple-douyin/pkg/constants"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if constants.DebugMode {
		r.Use(gin.Logger(), gin.Recovery())

	}
	unauthGroup := r.Group("/douyin")
	{
		unauthGroup.POST("/user/register/", user.Create)
		unauthGroup.POST("/user/login/", user.Check)
		unauthGroup.GET("/feed", video.Feed)
	}

	authGroup := r.Group("/douyin")
	authGroup.Use(auth.JWT())
	{
		authGroup.GET("/user/", user.Get)
		authGroup.POST("/publish/action/", video.Upload)
		authGroup.GET("/publish/list/", video.List)
		authGroup.POST("/favorite/action/", video.LikeOperation)
		authGroup.GET("/favorite/list/", video.LikeList)
		authGroup.POST("/comment/action/", video.CommentOperation)
		authGroup.GET("/comment/list/", video.CommentList)
		authGroup.POST("/relation/action/", user.FollowOperation)
		authGroup.GET("/relation/follow/list/", user.FollowList)
		authGroup.GET("/relation/follower/list/", user.FanList)
	}
	return r
}
