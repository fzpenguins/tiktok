package main

import (
	"testing"
	"tiktok/cmd/interaction/dal"
	"tiktok/cmd/interaction/rpc"
	"tiktok/cmd/interaction/service/comment"
	"tiktok/config"
	"tiktok/pkg/utils"
)

var (
	username           string
	uid                string
	vid                string
	cid                string
	token              string
	uidInt             int64
	interactionService *comment.CommentService
	pageNum            int64
	pageSize           int64
)

func TestMain(m *testing.M) {
	config.ConfigForTest()
	dal.Init()
	rpc.Init()

	username = "penqee"
	uid = "1"
	vid = "1"
	cid = "1"
	uidInt = 1
	pageNum = 0
	pageSize = 5
	token, _, _ = utils.GenerateToken(uidInt, username)

	m.Run()
}

func TestMainOrder(t *testing.T) {

	t.Run("create_comment", testCreateComment)

	t.Run("deelete_comment", testDeleteComment)

	t.Run("delete_comment", testDeleteComment)

	t.Run("RPC Test", testRPC)
}
