package service

import (
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/kitex_gen/video"
)

func (s *VideoService) Info(req *video.InfoReq) (resp []*db.Video, err error) {
	resp, err = dao.NewVideoDao(s.ctx).FindVideoByVid(req.Vid)
	if err != nil {
		return nil, err
	}
	for _, d := range resp {
		d.VisitCount = cache.VisitCount(s.ctx, d.Vid)
	}
	return resp, nil
}
