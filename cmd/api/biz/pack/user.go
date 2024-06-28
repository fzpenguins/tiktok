package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/user"
)

func BuildUser(user *user.User) *api.User {
	return &api.User{
		UID:       user.Uid,
		Username:  user.Username,
		AvatarURL: user.AvatarUrl,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
