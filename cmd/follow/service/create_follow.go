package service

import (
	"strconv"
	"tiktok/cmd/follow/dal/cache"
	"tiktok/cmd/follow/dal/db"
	"tiktok/cmd/follow/dal/db/dao"
	"tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"
	"time"

	"github.com/pkg/errors"
)

func (s *FollowService) CreateFollow(req *follow.ActionReq) error {

	uid, _ := strconv.ParseInt(req.Uid, 10, 64)
	toUid, _ := strconv.ParseInt(req.ToUid, 10, 64)
	if uid == toUid {
		return errno.FollowSelfError
	}
	if dao.NewFollowDao(s.ctx).IsFollowed(uid, toUid) {
		return errno.FollowExistedError
	}

	f := &db.Follow{
		FromUid:   uid,
		ToUid:     toUid,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err := cache.AddFollow(s.ctx, req.Uid, req.ToUid)
	if err != nil {
		return errors.WithMessage(err, errno.SetInfoError)
	}
	return dao.NewFollowDao(s.ctx).CreateFollow(f)
}
