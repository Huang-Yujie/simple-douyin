package oss

import (
	"io"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
)

var vodClient *vod.Client

type Video struct {
	Title    string
	Filename string
	File     io.Reader
}
