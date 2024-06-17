package main

import (
	"testing"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/utils"
)

func testDeleteFollow(t *testing.T) {
	cliam, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	err = followService.DeleteFollow(&follow.ActionReq{
		Uid:   cliam.Id,
		ToUid: uid2,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
