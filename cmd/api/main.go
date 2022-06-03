// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"net/http"
	constants2 "simple-douyin/pkg/constants"
	tracer2 "simple-douyin/pkg/tracer"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"simple-douyin/cmd/api/handlers"
	"simple-douyin/cmd/api/rpc"
)

func Init() {
	tracer2.InitJaeger(constants2.ApiServiceName)
	rpc.InitVideoRpc()
}

func main() {
	Init()
	r := gin.New()

	v1 := r.Group("/test")
	video1 := v1.Group("/video")
	video1.POST("", handlers.CreateVideo)
	video1.GET("/query", handlers.QueryVideoByUid)
	video1.GET("/queryByTime", handlers.QueryVideoByTime)
	video1.POST("/like", handlers.LikeVideo)
	video1.DELETE("/unlike", handlers.UnlikeVideo)
	video1.POST("/creatcomment", handlers.CreateComment)
	video1.DELETE("deletecomment", handlers.DeleteComment)
	video1.GET("/getcomment", handlers.GetComments)

	if err := http.ListenAndServe(":8088", r); err != nil {
		klog.Fatal(err)
	}
}
