package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/picture"
)

func BuildPicture(pic *picture.Image) *api.Image {
	return &api.Image{
		Pid: pic.Pid,
		URL: pic.Url,
	}
}

func BuildPictures(pics []*picture.Image) []*api.Image {
	if len(pics) == 0{
		return nil
	}
	images := make([]*api.Image, len(pics))
	for i, pic := range pics {
		images[i] = BuildPicture(pic)
	}
	return images
}