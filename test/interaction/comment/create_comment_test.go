package main

import (
	"strconv"
	"testing"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/utils"
)

func testCreateComment(t *testing.T) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	err = interactionService.CreateVideoComment(&interaction.PublishCommentReq{
		Vid:     vid,
		Uid:     strconv.FormatInt(claim.Uid, 10),
		Content: "content",
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
