package service

import (
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"

	"github.com/pkg/errors"
)

func (s *VideoService) GetPublishList(req *video.ListReq) (list []*video.Video, cnt int64, err error) {
	videos, cnt, err := dao.NewVideoDao(s.ctx).GetPublishList(req)
	if err != nil {
		return nil, 0, errors.WithMessage(err, errno.GetInfoError)
	}
	if videos == nil {
		return nil, 0, nil
	}
	list = make([]*video.Video, len(videos))
	for i, item := range videos {
		videoInfo, err := s.GetVideoInfo(s.ctx, item.Vid)
		if err != nil {
			return nil, 0, errors.WithMessage(err, errno.GetInfoError)
		}
		list[i] = videoInfo

	}

	return list, cnt, nil
}
