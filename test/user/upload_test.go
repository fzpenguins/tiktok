package main

import (
	"strconv"
	"testing"
	"tiktok/kitex_gen/user"
)

func testUpload(t *testing.T) {
	uidInt, _ := strconv.ParseInt(uid, 10, 64)
	_, err := userService.UploadAvatar(&user.UploadAvatarUrlReq{
		AvatarUrl: "",
		Uid:       uidInt,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
}
