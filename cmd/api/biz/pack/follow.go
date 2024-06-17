package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/follow"
)

func BuildFollow(f *follow.UserInfo) *api.UserInfo {
	return &api.UserInfo{
		UID:       f.Uid,
		Username:  f.Username,
		AvatarURL: f.AvatarUrl,
	}
}

func BuildFollows(f []*follow.UserInfo) []*api.UserInfo {
	if len(f) == 0 {
		return nil
	}
	resp := make([]*api.UserInfo, len(f))
	for _, info := range f {
		resp = append(resp, BuildFollow(info))
	}
	return resp
}
