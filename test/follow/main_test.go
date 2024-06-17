package main

import (
	"context"
	"testing"
	"tiktok/cmd/follow/dal"
	"tiktok/cmd/follow/rpc"
	"tiktok/cmd/follow/service"
	"tiktok/config"
	"tiktok/pkg/utils"
)

var (
	username string
	uid      string
	uidInt   int64
	token    string
	uid2     string

	pageSize int64
	pageNum  int64

	followService *service.FollowService
)

func TestMain(m *testing.M) {
	config.ConfigForTest()
	dal.Init()
	rpc.Init()

	username = "penqee"
	uid = "1"
	uid2 = "2"
	uidInt = 1
	pageNum = 0
	pageSize = 5
	followService = service.NewFollowService(context.Background())
	token, _, _ = utils.GenerateToken(uidInt, username)

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("follow", testCreateFollow)

	t.Run("unfollow", testDeleteFollow)

	t.Run("friend_list", testFriendList)

	t.Run("follow_list", testFollowList)

	t.Run("follower_list", testFollowerList)

	t.Run("RPC Test", testRPC)
}
