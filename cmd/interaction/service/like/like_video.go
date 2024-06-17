package like

import (
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
	"time"
)

func (s *LikeService) LikeVideo(req *interaction.ActionLikeReq) error {
	uid, _ := strconv.ParseInt(req.Uid, 10, 64)
	vid, _ := strconv.ParseInt(req.Vid, 10, 64)
	if !dao.NewLikeDao(s.ctx).GetUserLikeVideoExist(uid, vid) {
		return errno.LikeExistedError
	}
	l := &db.Like{
		Vid:       vid,
		Uid:       uid,
		Cid:       0,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	cache.AddLikeCount(s.ctx, uid, vid)

	return dao.NewLikeDao(s.ctx).CreateLike(l)
}
