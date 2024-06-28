package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tiktok/config"
	"tiktok/kitex_gen/picture"
	"tiktok/kitex_gen/picture/pictureservice"
	"tiktok/pkg/constants"
	"tiktok/pkg/errno"
	"tiktok/proto"
)

func InitConvertRPC() {
	conn, err := grpc.Dial("0.0.0.0:10010", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	pictureConClient = proto.NewPictureConServiceClient(conn)
	//convertClient = p.NewClipServiceClient(conn)
}

//func Insert(ctx context.Context, req *proto.ImageRequest) {
//	pictureConClient.GetImageVector()
//}

func InitPictureRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}

	c, err := pictureservice.NewClient(
		constants.PictureServiceName,
		client.WithMuxConnection(constants.MuxConnection),
		client.WithRPCTimeout(constants.RPCTimeout),
		client.WithConnectTimeout(constants.ConnectTimeout),
		client.WithResolver(r),
	)

	if err != nil {
		panic(err)
	}
	pictureClient = c
}

func InsertImage(ctx context.Context, req *picture.InsertRequest) (*picture.InsertResponse, error) {
	resp, err := pictureClient.Insert(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.picture Insert failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}

func SearchImage(ctx context.Context, req *picture.SearchByImageRequest) (*picture.SearchResponse, error) {
	resp, err := pictureClient.SearchByImage(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "api.rpc.picture Search failed")
	}

	if resp.Base.Code != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.Base.Code, resp.Base.Msg)
	}
	return resp, nil
}
