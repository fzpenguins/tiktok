package service

import (
	"tiktok/cmd/video/dal/db"
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/kitex_gen/video"
	"time"
)

func (s *VideoService) Feed(req *video.FeedReq) ([]*db.Video, error) {
	current := time.Unix(req.GetTime(), 0).Format("2006-01-02 15:04:05")
	return dao.NewVideoDao(s.ctx).FindVideosByTimeStr(current)
}
