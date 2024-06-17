package main

import (
	"testing"
	"tiktok/kitex_gen/interaction"
)

func testCommentList(t *testing.T) {
	_, err := interactionService.VideoCommentList(&interaction.ListCommentReq{
		Vid:      vid,
		PageSize: pageSize,
		PageNum:  pageNum,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
