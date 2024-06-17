package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/follow/followservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/errno"
)

func InitFollowRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := followservice.NewClient(
		constants.FollowServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	followClient = c
}

func FollowAction(ctx context.Context, req *follow.ActionReq) (*follow.ActionResp, error) {
	resp, err := followClient.Action(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.follow FollowAction failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil

}

func FollowingList(ctx context.Context, req *follow.ListFollowingReq) (*follow.ListFollowingResp, error) {
	resp, err := followClient.ListFollowing(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.follow FollowingList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *follow.ListFollowerReq) (*follow.ListFollowerResp, error) {
	resp, err := followClient.ListFollower(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.follow FollowerList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *follow.ListFriendReq) (*follow.ListFriendResp, error) {
	resp, err := followClient.ListFriend(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.follow FriendList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}
