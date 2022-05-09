package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	handler "simple-douyin/db_srv/handler/user"
	"simple-douyin/db_srv/initialize"
	proto "simple-douyin/db_srv/proto/user"
)

var userClient proto.UserClient
var conn *grpc.ClientConn

func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	userClient = proto.NewUserClient(conn)
}

func TestCreateUser() {
	// 初始化配置及其数据库连接
	resp, err := userClient.CreateUser(context.Background(), &proto.CreateUserReq{
		Mobile: "",
		Password: "",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
	return
}

func main() {
	initialize.InitConfig()
	initialize.InitDB()
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.Server{})
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:50051"))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()
	Init()
	TestCreateUser()
}
