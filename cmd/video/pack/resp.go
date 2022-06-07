package pack

import (
	"simple-douyin/kitex_gen/videoproto"
	"simple-douyin/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *videoproto.BaseResp {
	return baseResp(errno.ConvertErr(err))
}

func baseResp(err errno.ErrNo) *videoproto.BaseResp {
	return &videoproto.BaseResp{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
