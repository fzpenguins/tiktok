// Code generated by Kitex v0.9.1. DO NOT EDIT.

package followservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	follow "tiktok/kitex_gen/follow"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Action(ctx context.Context, req *follow.ActionReq, callOptions ...callopt.Option) (r *follow.ActionResp, err error)
	ListFollowing(ctx context.Context, req *follow.ListFollowingReq, callOptions ...callopt.Option) (r *follow.ListFollowingResp, err error)
	ListFollower(ctx context.Context, req *follow.ListFollowerReq, callOptions ...callopt.Option) (r *follow.ListFollowerResp, err error)
	ListFriend(ctx context.Context, req *follow.ListFriendReq, callOptions ...callopt.Option) (r *follow.ListFriendResp, err error)
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
	return &kFollowServiceClient{
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

type kFollowServiceClient struct {
	*kClient
}

func (p *kFollowServiceClient) Action(ctx context.Context, req *follow.ActionReq, callOptions ...callopt.Option) (r *follow.ActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Action(ctx, req)
}

func (p *kFollowServiceClient) ListFollowing(ctx context.Context, req *follow.ListFollowingReq, callOptions ...callopt.Option) (r *follow.ListFollowingResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFollowing(ctx, req)
}

func (p *kFollowServiceClient) ListFollower(ctx context.Context, req *follow.ListFollowerReq, callOptions ...callopt.Option) (r *follow.ListFollowerResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFollower(ctx, req)
}

func (p *kFollowServiceClient) ListFriend(ctx context.Context, req *follow.ListFriendReq, callOptions ...callopt.Option) (r *follow.ListFriendResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListFriend(ctx, req)
}
