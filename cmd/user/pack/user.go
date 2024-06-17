package pack

import (
	"tiktok/cmd/user/dal/db"
	"tiktok/kitex_gen/user"
)

func UserResp(data *db.User) *user.User {
	if data == nil {
		return nil
	}

	return &user.User{
		Uid:       data.Uid,
		Username:  data.Username,
		AvatarUrl: data.AvatarUrl,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func MFAResp(secret, code string) *user.QRCode {
	return &user.QRCode{
		Secret: secret,
		Qrcode: code,
	}
}
