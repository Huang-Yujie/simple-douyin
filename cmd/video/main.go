package main

import (
	"log"
	videoproto "simple-douyin/kitex_gen/videoproto/videoservice"
)

func main() {
	svr := videoproto.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
