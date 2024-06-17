package main

import (
	"testing"
	"tiktok/kitex_gen/follow/followservice"

	"github.com/cloudwego/kitex/client"
)

func testRPC(t *testing.T) {

	_, err := followservice.NewClient("follow",
		client.WithHostPorts("0.0.0.0:10004"),
	)

	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
