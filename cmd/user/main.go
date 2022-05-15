package main

import (
	"log"
	"simple-douyin/kitex_gen/userproto/userservice"
)

func main() {
	svr := userservice.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
