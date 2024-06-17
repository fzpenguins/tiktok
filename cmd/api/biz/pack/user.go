package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/user"
)

//var userKey UserInfo
//
//type UserInfo struct {
//	ID       int64  `json:"id"`
//	UserName string `json:"user_name"`
//}
//
//func GetUserInfo(ctx context.Context) (*UserInfo, error) {
//	user, ok := FromContext(ctx)
//	if !ok {
//		return nil, errors.New("获取用户信息错误")
//	}
//	return user, nil
//}
//
//func NewContext(ctx context.Context, u *UserInfo) context.Context {
//	return context.WithValue(ctx, userKey, u)
//}
//
//func FromContext(ctx context.Context) (*UserInfo, bool) {
//	u, ok := ctx.Value(userKey).(*UserInfo)
//	return u, ok
//}

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
