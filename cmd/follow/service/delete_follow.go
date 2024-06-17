package service

import (
	"github.com/pkg/errors"
	"tiktok/cmd/follow/dal/cache"
	"tiktok/cmd/follow/dal/db/dao"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"
)

func (s *FollowService) DeleteFollow(req *follow.ActionReq) error {
	err := cache.DeleteFollow(s.ctx, req.Uid, req.ToUid)
	if err != nil {
		return errors.WithMessage(err, errno.DeleteError)
	}

	err = dao.NewFollowDao(s.ctx).DeleteFollow(req.Uid, req.ToUid)
	if err != nil {
		return errors.WithMessage(err, errno.DeleteError)
	}

	return err
}
