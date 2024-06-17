package main

import (
	"testing"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/utils"
)

func testFollowerList(t *testing.T) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	_, err = followService.FollowerList(&follow.ListFollowerReq{
		Uid:      claim.Id,
		PageNum:  &pageNum,
		PageSize: &pageSize,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
