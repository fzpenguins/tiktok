package main

import (
	"testing"
	"tiktok/kitex_gen/user"
)

func testSearch(t *testing.T) {
	_, err := userService.SearchUserInfo(&user.InfoReq{
		Uid: &uid,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
