package pack

import (
	"tiktok/cmd/picture/dal/db"
	"tiktok/kitex_gen/picture"
)

func BuildImage(pic *db.Image) *picture.Image {
	return &picture.Image{
		Pid: pic.Pid,
		Url: pic.Url,
	}
}

func BuildImages(pics []*db.Image) []*picture.Image {
	if len(pics) == 0{
		return nil
	}
	images := make([]*picture.Image, len(pics))
	for i, pic := range pics {
		images[i] = BuildImage(pic)
	}
	return images
}
