package like

import (
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
)

func (s *LikeService) GetVideoInfo(req *interaction.GetVideoInfoRequest) (resp *interaction.GetVideoInfoResponse, err error) {
	resp = new(interaction.GetVideoInfoResponse)
	vid, _ := strconv.ParseInt(req.GetVid(), 10, 64)
	resp.CommentCount, err = cache.GetVideoCommentCount(s.ctx, req.GetVid())
	if err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	resp.LikeCount = cache.LikeCount(s.ctx, vid)
	return
}
