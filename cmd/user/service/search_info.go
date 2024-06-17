package service

import (
	"github.com/pkg/errors"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

func (s *UserService) SearchUserInfo(req *user.InfoReq) (*db.User, error) {
	usr, err := dao.NewUserDao(s.ctx).FindUserByUid(req.GetUid())
	if err != nil {
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}
	return usr, err
}
