package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"log"
	"tiktok/config"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/errno"
)

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	userClient = c
}

func Register(ctx context.Context, req *user.RegisterReq) (err error) {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return errors.WithMessage(err, "api.rpc.user Register failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}

func Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	resp, err = userClient.Login(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.user Login failed")
	}
	log.Println(22)
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}

func SearchInfo(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	resp, err = userClient.Info(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.user SearchInfo failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}

func UploadAvatar(ctx context.Context, req *user.UploadAvatarUrlReq) (resp *user.UploadAvatarUrlResp, err error) {
	resp, err = userClient.Upload(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.user UploadAvatar failed")
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}

func GetMFAqrcode(ctx context.Context, req *user.GetMFAReq) (resp *user.GetMFAResp, err error) {
	resp, err = userClient.GetMFA(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.user GetMFAqrcode failed")
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}

func MFABind(ctx context.Context, req *user.BindMFAReq) (resp *user.BindMFAResp, err error) {
	resp, err = userClient.BindMFA(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.user MFABind failed")
	}
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return
}
