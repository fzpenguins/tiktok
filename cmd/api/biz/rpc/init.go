package rpc

import (
	"tiktok/kitex_gen/follow/followservice"
	"tiktok/kitex_gen/interaction/interactionservice"
	picture "tiktok/kitex_gen/picture/pictureservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/kitex_gen/video/videoservice"
	"tiktok/proto"
)

var (
	userClient        userservice.Client
	videoClient       videoservice.Client
	interactionClient interactionservice.Client
	followClient      followservice.Client
	pictureClient     picture.Client
)

var (
	pictureConClient proto.PictureConServiceClient
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitFollowRPC()
	InitInteractionRPC()
	InitPictureRPC()
	// InitConvertRPC()
}
