package main

import (
	"context"
	"log"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/pack"
	"tiktok/cmd/user/service"
	user "tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
	"tiktok/pkg/utils"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterResp)

	if len(req.GetUsername()) == 0 || len(req.GetUsername()) > 10 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, errors.WithMessage(errno.ParamError, "Username length should be less than 10")
	}

	if len(req.GetPassword()) <= 5 {
		resp.Base = pack.BuildBaseResp(errno.ParamError)
		return resp, errors.WithMessage(errno.ParamError, "Password length should not be less than 5")
	}

	_, err = service.NewUserService(ctx).Register(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to register")
	}
	resp.Base = pack.BuildBaseResp(nil)
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginReq) (resp *user.LoginResp, err error) {
	// TODO: Your code here...
	resp = new(user.LoginResp)
	var usr *db.User
	usr, err = service.NewUserService(ctx).Login(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to login")
	}

	access_token, refresh_token, err := utils.GenerateToken(usr.Uid, req.Username)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)

		return resp, errors.WithMessage(err, errno.CreateFailed)
	}

	resp.Base = pack.BuildBaseResp(nil)

	resp.Data = pack.UserResp(usr)

	log.Println(resp)
	resp.Tokens = &user.Tokens{
		RefreshToken: refresh_token,
		AccessToken:  access_token,
	}

	return resp, nil
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoReq) (resp *user.InfoResp, err error) {
	// TODO: Your code here...
	resp = new(user.InfoResp)

	usr, err := service.NewUserService(ctx).SearchUserInfo(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to search info")
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.UserResp(usr)

	return
}

// Upload implements the UserServiceImpl interface.
func (s *UserServiceImpl) Upload(ctx context.Context, req *user.UploadAvatarUrlReq) (resp *user.UploadAvatarUrlResp, err error) {
	// TODO: Your code here...
	resp = new(user.UploadAvatarUrlResp)

	usr, err := service.NewUserService(ctx).UploadAvatar(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to upload avatar")
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.UserResp(usr)

	return
}

// GetMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetMFA(ctx context.Context, req *user.GetMFAReq) (resp *user.GetMFAResp, err error) {
	// TODO: Your code here...
	resp = new(user.GetMFAResp)

	usr, imgBase64, err := service.NewUserService(ctx).GetMfa(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to get mfa")
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = pack.MFAResp(usr.Secret, imgBase64)
	return
}

// BindMFA implements the UserServiceImpl interface.
func (s *UserServiceImpl) BindMFA(ctx context.Context, req *user.BindMFAReq) (resp *user.BindMFAResp, err error) {
	// TODO: Your code here...
	resp = new(user.BindMFAResp)

	err = service.NewUserService(ctx).BindMFA(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to bind mfa")
	}

	resp.Base = pack.BuildBaseResp(nil)

	return
}
