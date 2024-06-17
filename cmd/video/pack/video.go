package pack

import (
	"tiktok/cmd/video/dal/db"
	"tiktok/kitex_gen/video"
)

func BuildVideo(item *db.Video) *video.Video {
	return &video.Video{
		Vid:          item.Vid,
		Uid:          item.Uid,
		VideoUrl:     item.VideoUrl,
		CoverUrl:     item.CoverUrl,
		Title:        item.Title,
		Description:  item.Description,
		VisitCount:   int64(item.VisitCount),
		LikeCount:    int64(item.LikeCount),
		CommentCount: int64(item.CommentCount),
		CreatedAt:    item.CreatedAt,
		UpdatedAt:    item.UpdatedAt,
		DeletedAt:    item.DeletedAt,
	}
}

func BuildVideos(items []*db.Video) []*video.Video {
	if len(items) == 0 {
		return nil
	}
	// var ret []*video.Video
	ret := make([]*video.Video, len(items))
	for i, item := range items {
		ret[i] = BuildVideo(item)
	}
	// for _, item := range items {
	// 	ret = append(ret, BuildVideo(item))
	// }

	return ret
}
