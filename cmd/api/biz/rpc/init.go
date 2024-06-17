package rpc

import (
	"tiktok/kitex_gen/follow/followservice"
	"tiktok/kitex_gen/interaction/interactionservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/kitex_gen/video/videoservice"
)

var (
	userClient        userservice.Client
	videoClient       videoservice.Client
	interactionClient interactionservice.Client
	followClient      followservice.Client
)

func Init() {
	InitUserRPC()
	InitVideoRPC()
	InitFollowRPC()
	InitInteractionRPC()
}
