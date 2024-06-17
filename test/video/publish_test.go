package main

import (
	"testing"
	"tiktok/kitex_gen/video"

	"golang.org/x/sync/errgroup"
)

func testPublish(t *testing.T) {
	var err error
	var eg errgroup.Group
	var videoUrl, coverUrl string
	eg.Go(func() error {
		coverUrl, err = videoService.UploadCover(
			"https://th.bing.com/th/id/OIP.VlXsxUWAoGSSgksl1PTANwHaHa?rs=1&pid=ImgDetMain",
		)
		if err != nil {
			t.Logf("err: [%v] \n", err)
			t.Error(err)
			t.Fail()
		}
		return nil
	})

	eg.Go(
		func() error {
			videoUrl, err = videoService.UploadVideo(
				"https://th.bing.com/th/id/OIP.VlXsxUWAoGSSgksl1PTANwHaHa?rs=1&pid=ImgDetMain",
			)
			if err != nil {
				t.Logf("err: [%v] \n", err)
				t.Error(err)
				t.Fail()
			}
			return nil
		})

	if err = eg.Wait(); err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}
	_, err = videoService.CreateVideo(
		&video.PublishReq{
			Title:       "title",
			Description: "description",
			Uid:         uidInt,
		},
		videoUrl,
		coverUrl,
	)
	if err != nil {
		t.Logf("err: [%v] \n", err)
		t.Error(err)
		t.Fail()
	}

}
