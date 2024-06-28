package rpc

import (
	"context"
	"tiktok/proto"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitConvertRPC() {
	conn, err := grpc.Dial("0.0.0.0:10010", grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithTimeout(5*time.Second))
	if err != nil {
		panic(err)
	}
	pictureConClient = proto.NewPictureConServiceClient(conn)
	//convertClient = p.NewClipServiceClient(conn)
}

func GetImageVector(ctx context.Context, req *proto.ImageRequest) (vector []float32, err error) {
	resp, err := pictureConClient.GetImageVector(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.GetImageVector failed")
	}
	return resp.Vector, nil
}
