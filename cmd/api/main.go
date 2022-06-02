package main

import (
	"net/http"

	"simple-douyin/cmd/api/auth"
	"simple-douyin/cmd/api/handlers"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/pkg/config"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/oss"
	"simple-douyin/pkg/tracer"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	config.Init()
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.Init()
	oss.InitVodClient()
	auth.Init()
}

func main() {
	Init()
	r := handlers.NewRouter()
	if err := http.ListenAndServe(config.Server.HttpPort, r); err != nil {
		klog.Fatal(err)
	}
}
