package like

import (
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
)

func (s *LikeService) DeleteCommentLike(req *interaction.ActionLikeReq) error {
	uid, _ := strconv.ParseInt(req.GetUid(), 10, 64)
	cid, _ := strconv.ParseInt(req.GetVid(), 10, 64)
	cache.DecrCommentLikeCount(s.ctx, uid, cid)
	return dao.NewLikeDao(s.ctx).DeleteCommentLike(uid, cid)
}
