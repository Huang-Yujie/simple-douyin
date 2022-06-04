package main

import (
	"simple-douyin/cmd/user/dal"
	userproto "simple-douyin/kitex_gen/userproto/userservice"

	"net"

	"simple-douyin/pkg/config"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/middleware"
	tracer2 "simple-douyin/pkg/tracer"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	config.Init()
	tracer2.InitJaeger(constants.VideoServiceName)
	dal.Init()
}

func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.Server.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", config.Server.UserServiceAddr)
	if err != nil {
		panic(err)
	}
	svr := userproto.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
