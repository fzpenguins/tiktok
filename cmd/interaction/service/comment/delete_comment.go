package comment

import (
	"github.com/pkg/errors"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
)

func (s *CommentService) DeleteComment(req *interaction.DeleteReq) error {
	com, err := dao.NewCommentDao(s.ctx).DeleteComment(req)
	if err != nil {
		return errors.WithMessage(err, errno.DeleteError)
	}
	err = cache.DecrVideoCommentCount(s.ctx, req.GetVid(), req.Cid)
	if err != nil {
		return errors.WithMessage(err, errno.DeleteError)
	}
	if com.ParentId != 0 {
		err = cache.DecrChildCommentCount(s.ctx, strconv.FormatInt(com.ParentId, 10), strconv.FormatInt(com.Cid, 10))
	}
	return err
}
