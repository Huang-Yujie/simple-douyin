package main

import (
	"net/http"

	"simple-douyin/cmd/api/handlers"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/tracer"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Init() {
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.Init()
}

func main() {
	Init()
	r := handlers.NewRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
