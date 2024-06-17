package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"tiktok/config"
	"tiktok/kitex_gen/follow"
	"tiktok/kitex_gen/user"
	"tiktok/kitex_gen/user/userservice"
)

var (
	userClient userservice.Client
)

func InitUserRPC() {
	r, err := etcd.NewEtcdResolver([]string{config.EtcdAddr})
	if err != nil {
		panic(err)
	}
	userClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func GetUserInfo(ctx context.Context, uid string) (*follow.UserInfo, error) {
	resp, err := userClient.Info(ctx, &user.InfoReq{Uid: &uid})
	if err != nil {
		return nil, err
	}
	ret := &follow.UserInfo{
		Uid:       uid,
		Username:  resp.Data.Username,
		AvatarUrl: resp.Data.AvatarUrl,
	}
	return ret, nil
}
