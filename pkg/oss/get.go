package oss

import (
	"simple-douyin/pkg/config"
	"strconv"

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
	request := vod.CreateListSnapshotsRequest()
	request.VideoId = videoId
	request.SnapshotType = "CoverSnapshot"
	request.AuthTimeout = strconv.Itoa(config.Server.Timeout)
	response, err := vodClient.ListSnapshots(request)
	if err != nil {
		return "", err
	}
	return response.MediaSnapshot.Snapshots.Snapshot[0].Url, nil
}
