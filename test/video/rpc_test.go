package main

import (
	"testing"
	"tiktok/kitex_gen/interaction/interactionservice"

	"github.com/cloudwego/kitex/client"
)

func testRPC(t *testing.T) {
	_, err := interactionservice.NewClient("video",
		client.WithHostPorts("0.0.0.0:10002"),
	)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
