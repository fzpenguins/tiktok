package service

import (
	"github.com/pkg/errors"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
	"time"
)

func (s *VideoService) CreateVideo(req *video.PublishReq, videoUrl, coverUrl string) (*db.Video, error) {

	v := &db.Video{
		Uid:         req.Uid,
		VideoUrl:    videoUrl,
		CoverUrl:    coverUrl,
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	v, err := dao.NewVideoDao(s.ctx).CreateVideo(v)
	if err != nil {
		return nil, errors.WithMessage(err, errno.CreateFailed)
	}

	err = cache.AddVisitCount(s.ctx, v)
	if err != nil {
		return nil, errors.WithMessage(err, errno.CreateFailed)
	}
	return v, nil
}
