package oss

import (
	"simple-douyin/pkg/config"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

func GetPlayURL(videoId string) (string, error) {
	request := vod.CreateGetPlayInfoRequest()
	request.VideoId = videoId
	request.AuthTimeout = requests.NewInteger(config.Server.Timeout)
	request.StreamType = "video"
	request.Definition = "LD"
	request.AcceptFormat = "JSON"
	response, err := vodClient.GetPlayInfo(request)
	if err != nil {
		return "", err
	}
	playInfo := response.PlayInfoList.PlayInfo[0]
	return playInfo.PlayURL, nil
}
