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

func (s *LikeService) LikeComment(req *interaction.ActionLikeReq) error {
	uid, _ := strconv.ParseInt(req.Uid, 10, 64)
	cid, _ := strconv.ParseInt(req.Cid, 10, 64)
	if !dao.NewLikeDao(s.ctx).GetUserLikeCommentExist(uid, cid) {
		return errno.LikeExistedError
	}

	l := &db.Like{
		Vid:       0,
		Uid:       uid,
		Cid:       cid,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	cache.AddCommentLikeCount(s.ctx, uid, cid)

	return dao.NewLikeDao(s.ctx).CreateLike(l)
}
