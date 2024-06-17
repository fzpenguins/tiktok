package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/pkg/errno"
)

func BadResp() *api.BaseResp {
	return &api.BaseResp{
		Code: errno.FailureCode,
		Msg:  errno.FailureMsg,
	}
}

func GoodResponse() *api.BaseResp {
	return &api.BaseResp{
		Code: errno.SuccessCode,
		Msg:  errno.SuccessMsg,
	}
}
