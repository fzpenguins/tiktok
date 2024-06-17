package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"tiktok/config"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/interaction/interactionservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/errno"
)

func InitInteractionRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := interactionservice.NewClient(
		constants.InteractionServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	interactionClient = c
}

func LikeAction(ctx context.Context, req *interaction.ActionLikeReq) (*interaction.ActionLikeResp, error) {
	resp, err := interactionClient.ActionLike(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.interaction LikeAction failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func LikeList(ctx context.Context, req *interaction.ListLikeReq) (*interaction.ListLikeResp, error) {
	resp, err := interactionClient.ListLike(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.interaction LikeList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func CommentAction(ctx context.Context, req *interaction.PublishCommentReq) (*interaction.PublishCommentResp, error) {
	resp, err := interactionClient.PublishComment(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.interaction CommentAction failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func CommentList(ctx context.Context, req *interaction.ListCommentReq) (*interaction.ListCommentResp, error) {
	resp, err := interactionClient.ListComment(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.interaction CommentList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func DeleteComment(ctx context.Context, req *interaction.DeleteReq) (*interaction.DeleteResp, error) {
	resp, err := interactionClient.Delete(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.interaction DeleteComment failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}
