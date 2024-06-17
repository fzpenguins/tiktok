package main

import (
	"testing"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/utils"
)

func testDeleteComment(t *testing.T) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	err = interactionService.DeleteComment(&interaction.DeleteReq{
		Cid: cid,
		Uid: claim.Id,
	})
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
