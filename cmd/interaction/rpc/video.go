package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/config"
	"tiktok/kitex_gen/video"
	"tiktok/kitex_gen/video/videoservice"
)

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}
	videoClient, err = videoservice.NewClient("video", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func GetVideoInfo(ctx context.Context, vid []int64) (*video.InfoResp, error) {
	return videoClient.Info(ctx, &video.InfoReq{Vid: vid})
}
