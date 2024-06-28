package comment

import (
	"strconv"
	"sync"
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
	"tiktok/cmd/interaction/dal/db/dao"
	"tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"

	"github.com/pkg/errors"
)

func (s *CommentService) VideoCommentList(req *interaction.ListCommentReq) ([]*interaction.Comment, error) {

	var list []*db.Comment

	var err error

	if cache.IsVideoCommentListExist(s.ctx, req.Vid) {
		rets := cache.GetVideoCommentList(s.ctx, req.GetVid(), req.GetPageSize(), req.GetPageNum())

		if len(rets) == 0 {
			list, err = dao.NewCommentDao(s.ctx).GetCommentsByVid(req)
			if err != nil {
				return nil, errors.WithMessage(err, errno.GetInfoError)
			}
		} else {
			return rets, nil
		}
	} else {
		list, err = dao.NewCommentDao(s.ctx).GetCommentsByVid(req)
		if err != nil {
			return nil, errors.WithMessage(err, errno.GetInfoError)
		}
	}
	comments := make([]*interaction.Comment, len(list))
	var wg sync.WaitGroup
	for i, item := range list {
		wg.Add(1)

		go func(item *db.Comment, index int) {
			defer wg.Done()
			t := &interaction.Comment{
				Uid:       strconv.FormatInt(item.Uid, 10),
				Vid:       strconv.FormatInt(item.Vid, 10),
				Cid:       strconv.FormatInt(item.Cid, 10),
				ParentId:  "0",
				Content:   item.Content,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
				DeletedAt: item.DeletedAt,
			}

			t.ChildCount, err = cache.GetChildCommentCount(s.ctx, strconv.FormatInt(item.Cid, 10))
			if err != nil {
				return
			}

			t.LikeCount = cache.CommentLikeCount(s.ctx, item.Cid)
			err = cache.AddIntoVideoCommentList(s.ctx, t, strconv.FormatInt(item.Vid, 10))
			if err != nil {
				return
			}

			comments[index] = t

		}(item, i)

	}

	wg.Wait()

	return comments, err

}
