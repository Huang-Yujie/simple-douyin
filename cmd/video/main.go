package main

import (
	videoproto "simple-douyin/kitex_gen/videoproto/videoservice"

	"net"
	"simple-douyin/cmd/video/rpc"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"simple-douyin/pkg/bound"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/middleware"
	tracer2 "simple-douyin/pkg/tracer"
)

func Init() {
	tracer2.InitJaeger(constants.VideoServiceName)
	rpc.InitRPC()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress}) // r should not be reused.
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	if err != nil {
		panic(err)
	}
	Init()
	svr := videoproto.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                             // middleWare
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
