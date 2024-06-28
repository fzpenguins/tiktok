package service

import (
	"strconv"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/cmd/video/rpc"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
	"time"

	"github.com/pkg/errors"
)

func (s *VideoService) Search(req *video.SearchReq) ([]*db.Video, int64, error) {

	queryDao := dao.NewVideoDao(s.ctx)
	queryDao.DB = queryDao.GetVideoByKeyword(req.Keywords)

	if req.GetUsername() != "" {
		uid, err := cache.GetUid(s.ctx, req)
		if err != nil {
			return nil, 0, errors.WithMessage(err, errno.ParseFailed)
		}
		queryDao.DB = queryDao.GetVideoByUid(uid)
	}

	if req.GetFromDate() > 0 {
		FromDateString := time.Unix(req.GetFromDate(), 0).Format("2006-01-02 15:04:05")
		queryDao.DB = queryDao.GetVideoByFromDate(FromDateString)
	}

	if req.GetToDate() > 0 {
		ToDateString := time.Unix(req.GetToDate(), 0).Format("2006-01-02 15:04:05")
		queryDao.DB = queryDao.Where("created_at >= ?", ToDateString)
	}
	var size int64
	var list []*db.Video
	var err error
	list, size, err = queryDao.GetVideos(req.GetPageNum(), req.GetPageSize())
	if err != nil {
		return nil, 0, errors.WithMessage(err, errno.GetInfoError)
	}

	for _, item := range list {
		item.VisitCount = cache.VisitCount(s.ctx, item.Vid)
		item.LikeCount, item.CommentCount, err = rpc.GetVideoInfo(s.ctx, &interaction.GetVideoInfoRequest{
			Vid: strconv.FormatInt(item.Vid, 10)},
		)
	}
	return list, size, err
}
