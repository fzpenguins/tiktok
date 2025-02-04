// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "tiktok/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error)
	Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error)
	Info(ctx context.Context, req *user.InfoReq, callOptions ...callopt.Option) (r *user.InfoResp, err error)
	Upload(ctx context.Context, req *user.UploadAvatarUrlReq, callOptions ...callopt.Option) (r *user.UploadAvatarUrlResp, err error)
	GetMFA(ctx context.Context, req *user.GetMFAReq, callOptions ...callopt.Option) (r *user.GetMFAResp, err error)
	BindMFA(ctx context.Context, req *user.BindMFAReq, callOptions ...callopt.Option) (r *user.BindMFAResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) Register(ctx context.Context, req *user.RegisterReq, callOptions ...callopt.Option) (r *user.RegisterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kUserServiceClient) Login(ctx context.Context, req *user.LoginReq, callOptions ...callopt.Option) (r *user.LoginResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kUserServiceClient) Info(ctx context.Context, req *user.InfoReq, callOptions ...callopt.Option) (r *user.InfoResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Info(ctx, req)
}

func (p *kUserServiceClient) Upload(ctx context.Context, req *user.UploadAvatarUrlReq, callOptions ...callopt.Option) (r *user.UploadAvatarUrlResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Upload(ctx, req)
}

func (p *kUserServiceClient) GetMFA(ctx context.Context, req *user.GetMFAReq, callOptions ...callopt.Option) (r *user.GetMFAResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetMFA(ctx, req)
}

func (p *kUserServiceClient) BindMFA(ctx context.Context, req *user.BindMFAReq, callOptions ...callopt.Option) (r *user.BindMFAResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.BindMFA(ctx, req)
}
