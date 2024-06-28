package service

import (
	"strconv"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"

	"github.com/pkg/errors"
)

func (s *UserService) UploadAvatar(req *user.UploadAvatarUrlReq) (*db.User, error) {
	usr, err := dao.NewUserDao(s.ctx).FindUserByUid(strconv.FormatInt(req.Uid, 10))
	if err != nil {
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	err = dao.NewUserDao(s.ctx).UploadAvatar(req.AvatarUrl, usr)
	if err != nil {
		return nil, errors.WithMessage(err, errno.SetInfoError)
	}

	return usr, err
}
