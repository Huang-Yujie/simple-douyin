package oss

import (
	"encoding/base64"
	"encoding/json"
	"io"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/vod"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func initOssClient(uploadAuthDTO UploadAuthDTO, uploadAddressDTO UploadAddressDTO) (*oss.Client, error) {
	client, err := oss.New(uploadAddressDTO.Endpoint,
		uploadAuthDTO.AccessKeyId,
		uploadAuthDTO.AccessKeySecret,
		oss.SecurityToken(uploadAuthDTO.SecurityToken),
		oss.Timeout(86400*7, 86400*7))
	return client, err
}

func uploadLocalFile(client *oss.Client, uploadAddressDTO UploadAddressDTO, localFile io.Reader) error {
	// 获取存储空间。
	bucket, err := client.Bucket(uploadAddressDTO.Bucket)
	if err != nil {
		return err
	}
	// 上传本地文件。
	err = bucket.PutObject(uploadAddressDTO.FileName, localFile)
	if err != nil {
		return err
	}
	return nil
}

type UploadAuthDTO struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}
type UploadAddressDTO struct {
	Endpoint string
	Bucket   string
	FileName string
}

func Upload(v *Video) (string, error) {
	// 获取上传地址和凭证
	request := vod.CreateCreateUploadVideoRequest()
	request.Title = v.Title
	request.FileName = v.Filename
	request.AcceptFormat = "JSON"
	response, err := vodClient.CreateUploadVideo(request)
	if err != nil {
		return "", err
	}
	// 执行成功会返回VideoId、UploadAddress和UploadAuth
	var uploadAuthDTO UploadAuthDTO
	var uploadAddressDTO UploadAddressDTO
	uploadAuthDecode, err := base64.StdEncoding.DecodeString(response.UploadAuth)
	if err != nil {
		return "", err
	}
	uploadAddressDecode, err := base64.StdEncoding.DecodeString(response.UploadAddress)
	if err != nil {
		return "", err
	}
	json.Unmarshal(uploadAuthDecode, &uploadAuthDTO)
	json.Unmarshal(uploadAddressDecode, &uploadAddressDTO)
	// 使用UploadAuth和UploadAddress初始化OSS客户端
	ossClient, err := initOssClient(uploadAuthDTO, uploadAddressDTO)
	if err != nil {
		return "", err
	}
	// 上传文件，注意是同步上传会阻塞等待，耗时与文件大小和网络上行带宽有关
	if err := uploadLocalFile(ossClient, uploadAddressDTO, v.File); err != nil {
		return "", err
	}
	return response.VideoId, nil
}
