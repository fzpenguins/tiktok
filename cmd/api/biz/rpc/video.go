package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"log"
	"tiktok/config"
	"tiktok/kitex_gen/video"
	"tiktok/kitex_gen/video/videoservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/errno"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	videoClient = c
}

func UploadVideo(ctx context.Context, req *video.PublishReq) (*video.PublishResp, error) {
	log.Println(22222)
	resp, err := videoClient.Publish(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video UploadVideo failed")
	}
	log.Println("....")
	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func PopularRankList(ctx context.Context, req *video.PopularReq) (*video.PopularResp, error) {
	resp, err := videoClient.Popular(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video PopularRankList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func PublishList(ctx context.Context, req *video.ListReq) (*video.ListResp, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video PublishList failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func SearchVideo(ctx context.Context, req *video.SearchReq) (*video.SearchResp, error) {
	resp, err := videoClient.Search(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video SearchVideo failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func Feed(ctx context.Context, req *video.FeedReq) (*video.FeedResp, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.video Feed failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}
