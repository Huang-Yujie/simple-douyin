package main

import (
	"log"
	"simple-douyin/cmd/video/test"
	"simple-douyin/kitex_gen/videoproto/videoservice"
)

func main() {
	svr := videoservice.NewServer(new(test.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
