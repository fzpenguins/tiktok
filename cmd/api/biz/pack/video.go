package pack

import (
	"tiktok/cmd/api/biz/model/api"
	"tiktok/kitex_gen/video"
)

func BuildVideo(v *video.Video) *api.Video {
	return &api.Video{
		Vid:          v.Vid,
		UID:          v.Uid,
		VideoURL:     v.VideoUrl,
		CoverURL:     v.CoverUrl,
		Title:        v.Title,
		Description:  v.Description,
		VisitCount:   v.VisitCount,
		LikeCount:    v.LikeCount,
		CommentCount: v.CommentCount,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
		DeletedAt:    v.DeletedAt,
	}
}

func BuildVideos(videos []*video.Video) []*api.Video {
	if len(videos) == 0 {
		return nil
	}
	resp := make([]*api.Video, len(videos))
	for _, v := range videos {
		resp = append(resp, BuildVideo(v))
	}

	return resp
}
