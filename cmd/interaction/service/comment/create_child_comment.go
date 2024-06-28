package comment

import (
	"strconv"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func (s *CommentService) CreateChildComment(req *interaction.PublishCommentReq) error {
	cid, _ := strconv.ParseInt(req.GetCid(), 10, 64)
	u, _ := strconv.ParseInt(req.GetUid(), 10, 64)

	parentC, err := dao.NewCommentDao(s.ctx).GetCommentByCid(cid)
	if err != nil {
		return errors.WithMessage(err, errno.GetInfoError)
	}

	c := &db.Comment{
		Uid:       u,
		Vid:       parentC.Vid,
		ParentId:  cid,
		Content:   req.Content,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = dao.NewCommentDao(s.ctx).CreateComment(c)
	if err != nil {
		return errors.WithMessage(err, errno.CreateFailed)
	}
	eg, ctx := errgroup.WithContext(s.ctx)
	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()
		err = cache.AddVideoCommentCount(s.ctx, strconv.FormatInt(parentC.Vid, 10), strconv.FormatInt(c.Cid, 10))
		return errors.WithMessage(err, errno.CreateFailed)
	})

	eg.Go(func() error {
		defer func() {
			if e := recover(); e != nil {
				klog.Error(e)
			}
		}()

		err := cache.AddChildCommentCount(ctx, req.Cid, strconv.FormatInt(c.Cid, 10))
		return errors.WithMessage(err, errno.CreateFailed)
	})

	if err := eg.Wait(); err != nil {
		return errors.WithMessage(err, errno.CreateFailed)
	}
	return nil
}
