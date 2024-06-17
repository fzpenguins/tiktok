package service

import (
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

//type UserAvatarUrl struct {
//	Ctx context.Context
//	Url string
//	Uer *db.User
//}

//var UserAvatarUrlChannel = make(chan *UserAvatarUrl)

func (s *UserService) UploadAvatar(req *user.UploadAvatarUrlReq) (*db.User, error) {
	usr, err := dao.NewUserDao(s.ctx).FindUserByUid(strconv.FormatInt(req.Uid, 10))
	if err != nil {
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	err = dao.NewUserDao(s.ctx).UploadAvatar(req.AvatarUrl, usr)
	if err != nil {
		return nil, errors.WithMessage(err, errno.SetInfoError)
	}
	//userReq := &UserAvatarUrl{
	//	Ctx: s.ctx,
	//	Url: req.AvatarUrl,
	//	Uer: usr,
	//}
	//UserAvatarUrlChannel <- userReq

	//usr.AvatarUrl = req.GetAvatarUrl()
	//usr.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	return usr, err
}

//func HandleUploadAvatarUrl() {
//	for {
//		select {
//		case info := <-UserAvatarUrlChannel:
//			userDao := dao.NewUserDao(info.Ctx)
//			err := userDao.UpdateDate(info.Uer)
//			if err != nil {
//				log.Println("failed to update date,err = ", err)
//			}
//
//			err = userDao.UploadAvatar(info.Url, info.Uer.Uid)
//			if err != nil {
//				log.Println("failed to upload avatar,err =", err)
//			}
//
//		default:
//
//		}
//	}
//}
