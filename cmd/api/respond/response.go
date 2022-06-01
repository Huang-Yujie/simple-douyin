package respond

type BaseResp struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

var Success = BaseResp{
	Code: 0,
	Msg:  "success",
}
