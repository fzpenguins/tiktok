package service

import (
	"github.com/pkg/errors"
	"log"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
)

func (s *UserService) Login(req *user.LoginReq) (*db.User, error) {
	usr, err := dao.NewUserDao(s.ctx).FindUserByName(req.Username)
	if err != nil {
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	if !usr.VerifyPassword(req.Password) {
		return nil, errors.WithMessage(errno.ParamError, "wrong password")
	}
	err = dao.NewUserDao(s.ctx).UpdateDate(usr)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to update")
	}
	log.Println(usr)
	return usr, err
}
