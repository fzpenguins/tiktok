package main

import (
	"testing"
	"tiktok/kitex_gen/video/videoservice"

	"github.com/cloudwego/kitex/client"
)

func testRPC(t *testing.T) {
	_, err := videoservice.NewClient("video",
		client.WithHostPorts("0.0.0.0:10003"),
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

}
