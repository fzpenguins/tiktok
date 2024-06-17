package main

import (
	"context"
	"strconv"
	"testing"
	"tiktok/cmd/video/dal"
	"tiktok/cmd/video/rpc"
	"tiktok/cmd/video/service"
	"tiktok/config"
	"tiktok/pkg/utils"
)

var (
	token    string
	uid      string
	username string
	uidInt   int64

	videoService *service.VideoService
	PageNum      int64
	PageSize     int64
)

func TestMain(m *testing.M) {
	config.ConfigForTest()
	dal.Init()
	rpc.Init()

	videoService = service.NewVideoService(context.Background())

	uid = "1"
	username = "penqee"
	uidInt, _ = strconv.ParseInt(uid, 10, 64)

	PageNum = 0
	PageSize = 5

	token, _, _ = utils.GenerateToken(uidInt, username)

	m.Run()
}

func TestMainOrder(t *testing.T) {

	t.Run("publish", testPublish)

	t.Run("search", testSearch)

	t.Run("publish_list", testPublishList)

	t.Run("rank_list", testRankList)

	t.Run("RPC Test", testRPC)
}
