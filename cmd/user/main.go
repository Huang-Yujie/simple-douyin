package main

import (
	"log"
	userproto "simple-douyin/kitex_gen/userproto/userservice"
)

func main() {
	svr := userproto.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
