package video

import (
	"simple-douyin/cmd/api/forms"
	"simple-douyin/cmd/api/respond"
	"simple-douyin/cmd/api/rpc"
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/constants"
	"simple-douyin/pkg/oss"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	appUserID := c.GetInt64(constants.IdentityKey)
	param := new(forms.VideoUploadReq)
	if err := c.ShouldBind(param); err != nil {
		respond.Error(c, err)
		return
	}
	fileHeader, err := c.FormFile("data")
	if err != nil {
		respond.Error(c, err)
	}
	file, err := fileHeader.Open()
	if err != nil {
		respond.Error(c, err)
	}
	defer file.Close()
	ossUploadReq := &oss.Video{
		Title:    param.Title,
		Filename: "/simple-douyin/" + fileHeader.Filename,
		File:     file,
	}
	ossVideoID, err := oss.Upload(ossUploadReq)
	if err != nil {
		respond.Error(c, err)
		return
	}
	req := &videoproto.CreateVideoReq{
		VideoBaseInfo: &videoproto.VideoBaseInfo{
			UserId:     appUserID,
			OssVideoId: ossVideoID,
			Title:      param.Title,
		},
	}
	if err := rpc.CreateVideo(c, req); err != nil {
		respond.Error(c, err)
		return
	}
	// if err := oss.Snapshot(ossVideoID); err != nil {
	// 	respond.Error(c, err)
	// 	return
	// }
	respond.OK(c)
}
