package like

import (
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
)

func (s *LikeService) DeleteVideoLike(req *interaction.ActionLikeReq) error {
	uid, _ := strconv.ParseInt(req.GetUid(), 10, 64)
	vid, _ := strconv.ParseInt(req.GetVid(), 10, 64)
	cache.DecrLikeCount(s.ctx, uid, vid)
	return dao.NewLikeDao(s.ctx).DeleteVideoLike(uid, vid)
}
