package main

import (
	"testing"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/utils"
)

func testCreateFollow(t *testing.T) {
	cliam, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	err = followService.CreateFollow(&follow.ActionReq{
		ToUid: uid2,
		Uid:   cliam.Id,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
