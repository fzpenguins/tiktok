package main

import (
	"testing"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/utils"
)

func testPublishList(t *testing.T) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	_, _, err = videoService.GetPublishList(&video.ListReq{
		PageNum:  PageNum,
		PageSize: PageSize,
		Uid:      claim.Id,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
