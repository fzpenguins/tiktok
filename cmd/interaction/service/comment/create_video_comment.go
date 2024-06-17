package comment

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
	"time"
)

func (s *CommentService) CreateVideoComment(req *interaction.PublishCommentReq) error {
	vid, _ := strconv.ParseInt(req.GetVid(), 10, 64)
	u, _ := strconv.ParseInt(req.GetUid(), 10, 64)
	c := &db.Comment{
		Uid:       u,
		Vid:       vid,
		Content:   req.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	eg, ctx := errgroup.WithContext(s.ctx)

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		err := dao.NewCommentDao(ctx).CreateComment(c)
		return errors.WithMessage(err, errno.CreateFailed)
	})

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()

		err := cache.AddVideoCommentCount(ctx, req.Vid, strconv.FormatInt(c.Cid, 10))
		return errors.WithMessage(err, errno.CreateFailed)
	})

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		if cache.IsVideoCommentListExist(ctx, req.Cid) {
			err := cache.DeleteAllItemInVideoCommentList(ctx, req.Cid)
			return errors.WithMessage(err, errno.DeleteError)
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		return errors.WithMessage(err, errno.CreateFailed)
	}
	return nil
}
