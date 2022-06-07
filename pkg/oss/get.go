package oss

import (
	"fmt"
	"simple-douyin/pkg/config"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

func GetPlayURL(videoId string) (string, error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AuthTimeout = requests.NewInteger(config.Server.Timeout)
	request.StreamType = "video"
	request.AcceptFormat = "JSON"
	response, err := vodClient.GetPlayInfo(request)
	if err != nil {
		return "", err
	}
	return response.PlayInfoList.PlayInfo[0].PlayURL, nil
}

func GetCoverURL(videoId string) (string, error) {
	request := vod.CreateGetVideoInfoRequest()
	request.VideoId = videoId
	response, err := vodClient.GetVideoInfo(request)
	if err != nil {
		return "", err
	}
	fmt.Println(response.Video)
	return response.Video.CoverURL, nil
}
