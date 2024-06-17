package main

import (
	"testing"
	"tiktok/kitex_gen/video"
)

func testSearch(t *testing.T) {
	_, _, err := videoService.Search(&video.SearchReq{
		Keywords: "key",
		PageSize: PageSize,
		PageNum:  PageNum,
		Username: &username,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
