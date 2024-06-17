package service

import (
	"tiktok/cmd/video/dal/cache"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"

	"github.com/pkg/errors"
)

func (s *VideoService) GetRankList(req *video.PopularReq) (list []*video.Video, err error) {
	vids, _, _, err := cache.GetVideoRank(s.ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	if vids == nil {
		return nil, nil
	}

	list = make([]*video.Video, len(vids))
	for i, v := range vids {
		videoInfo, err := s.GetVideoInfo(s.ctx, v)
		if err != nil {
			return nil, errors.WithMessage(err, errno.GetInfoError)
		}
		list[i] = videoInfo
	}

	return list, nil
}
