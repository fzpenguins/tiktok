package rpc

import "tiktok/proto"

var (
	pictureConClient proto.PictureConServiceClient
)

func Init() {
	InitConvertRPC()
}
