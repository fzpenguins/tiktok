package main

import (
	"testing"
	"tiktok/kitex_gen/user"
)

func testRegister(t *testing.T) {
	_, err := userService.Register(&user.RegisterReq{
		Username: username,
		Password: password,
	})

	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
