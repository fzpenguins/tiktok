package service

import (
	"context"
	"strconv"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db/dao"
	"tiktok/cmd/video/pack"
	"tiktok/cmd/video/rpc"
	"tiktok/kitex_gen/interaction"
	"tiktok/kitex_gen/video"
	"tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/common/json"
	"github.com/pkg/errors"
)

func (s *VideoService) GetVideoInfo(ctx context.Context, vid int64) (videos *video.Video, err error) {
	vidStr := strconv.FormatInt(vid, 10)
	videos = new(video.Video)
	if ok := cache.IsExistsVideoInfo(ctx, vidStr); ok {
		data, err := cache.GetVideoInfo(ctx, vidStr)
		if err != nil {
			return nil, errors.WithMessage(err, errno.GetInfoError)
		}
		err = json.Unmarshal([]byte(data), &videos)
		if err != nil {
			return nil, errors.Wrap(err, errno.ParseFailed)
		}

		return videos, nil
	}

	v, err := dao.NewVideoDao(ctx).GetVideoByVid(vid)
	if err != nil {
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}
	if v == nil {
		return nil, nil
	}

	//将新的数据写入redis
	errCh := make(chan error, 1)
	go func() {
		videoJson, err := json.Marshal(pack.BuildVideo(v))
		if err != nil {
			errCh <- err
			return
		}
		err = cache.AddVideoInfo(ctx, vidStr, videoJson)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	v.VisitCount = cache.VisitCount(s.ctx, vid)
	if err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	if err := <-errCh; err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	//rpc
	v.LikeCount, v.CommentCount, err = rpc.GetVideoInfo(ctx, &interaction.GetVideoInfoRequest{Vid: vidStr})
	if err != nil {
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}

	return pack.BuildVideo(v), nil
}
