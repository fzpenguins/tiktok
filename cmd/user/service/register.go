package service

import (
	"tiktok/cmd/user/dal/cache"
	"tiktok/cmd/user/dal/db"
	"tiktok/cmd/user/dal/db/dao"
	"tiktok/kitex_gen/user"
	"tiktok/pkg/errno"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *UserService) Register(req *user.RegisterReq) (*db.User, error) {
	userModel := &db.User{
		Username: req.Username,
		Password: req.Password,
	}

	_, err := dao.NewUserDao(s.ctx).FindUserByName(req.Username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = userModel.SetPassword(req.Password)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to set password")
		}
		userModel.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

		err = dao.NewUserDao(s.ctx).CreateUser(userModel)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create user")
		}

		err = cache.SetUsernameHashUid(s.ctx, req.GetUsername(), userModel.Uid)
		if err != nil {
			return nil, errors.WithMessage(err, "failed to create user info")
		}
		return userModel, nil
	}
	if err != nil {
		return nil, errors.WithMessage(err, "database error")
	}
	return nil, errno.UserExistedError
}
