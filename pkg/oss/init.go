package oss

import (
	"simple-douyin/pkg/config"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/cloudwego/kitex/pkg/klog"
)

func InitVodClient() {
	// 点播服务接入区域
	regionId := "cn-beijing"
	// 创建授权对象
	credential := &credentials.AccessKeyCredential{
		AccessKeyId:     config.OSS.KeyID,
		AccessKeySecret: config.OSS.KeySecret,
	}
	// 自定义config
	config := sdk.NewConfig()
	config.AutoRetry = true     // 失败是否自动重试
	config.MaxRetryTime = 3     // 最大重试次数
	config.Timeout = 3000000000 // 连接超时，单位：纳秒；默认为3秒
	// 创建vodClient实例
	var err error
	vodClient, err = vod.NewClientWithOptions(regionId, config, credential)
	if err != nil {
		klog.Fatal(err)
	}
}
