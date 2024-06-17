package main

import (
	"strconv"
	"testing"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/utils"
)

func testLogin(t *testing.T) {
	resp, err := userService.Login(&user.LoginReq{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	access_token, refresh_token, err = utils.GenerateToken(resp.Uid, resp.Username)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

	uid = strconv.FormatInt(resp.Uid, 10)
}
