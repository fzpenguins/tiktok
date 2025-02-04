package pack

import (
	"github.com/pkg/errors"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"
)

func BuildBaseResp(err error) *follow.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}

	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceError.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *follow.BaseResp {
	return &follow.BaseResp{
		Code: err.ErrorCode,
		Msg:  err.ErrorMsg,
	}
}
