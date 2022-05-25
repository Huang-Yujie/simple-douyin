package main

import (
	"log"
	videoproto "simple-douyin/cmd/video/kitex_gen/videoproto/videoservice"
	"simple-douyin/cmd/video/test"
)

func main() {
	svr := videoproto.NewServer(new(test.VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
