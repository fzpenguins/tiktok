package main

import (
	"testing"
	"tiktok/kitex_gen/video"
)

func testRankList(t *testing.T) {

	_, err := videoService.GetRankList(&video.PopularReq{
		PageSize: &PageSize,
		PageNum:  &PageNum,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
