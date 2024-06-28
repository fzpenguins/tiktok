package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"golang.org/x/sync/errgroup"

	"strconv"
	"sync"
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/pack"
	"tiktok/cmd/video/rpc"
	"tiktok/cmd/video/service"
	"tiktok/kitex_gen/interaction"
	video "tiktok/kitex_gen/video"
	"tiktok/pkg/errno"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedReq) (resp *video.FeedResp, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResp)

	if req.Time == nil {
		current := time.Now().UnixMilli()
		req.Time = &current
	}

	videoList, err := service.NewVideoService(ctx).Feed(req)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	var wg sync.WaitGroup
	for _, v := range videoList {
		wg.Add(1)
		go func() {
			v.VisitCount = cache.VisitCount(ctx, v.Vid)
			v.LikeCount, v.CommentCount, _ = rpc.GetVideoInfo(ctx, &interaction.GetVideoInfoRequest{Vid: strconv.FormatInt(v.Vid, 10)})
		}()
		wg.Done()
	}
	wg.Wait()
	resp.Items = pack.BuildVideos(videoList)
	return
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishReq) (resp *video.PublishResp, err error) {
	// TODO: Your code here...

	resp = new(video.PublishResp)
	var eg errgroup.Group
	var videoUrl, coverUrl string

	eg.Go(func() error {
		videoUrl, err = service.NewVideoService(ctx).UploadVideo(req.Data.VideoUrl)
		if err != nil {
			klog.Error(err)
			return errno.FileUploadError
		}
		return nil
	})
	// 截取并上传封面
	eg.Go(func() error {
		coverUrl, err = service.NewVideoService(ctx).UploadCover(req.Data.CoverUrl)
		if err != nil {
			klog.Error(err)
			return errno.FileUploadError
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	_, err = service.NewVideoService(ctx).CreateVideo(req, videoUrl, coverUrl)
	if err != nil {
		klog.Error(err)
		return nil, err
	}

	resp.Base = pack.BuildBaseResp(nil)

	return resp, err
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, req *video.ListReq) (resp *video.ListResp, err error) {
	// TODO: Your code here...
	resp = new(video.ListResp)
	list, count, err := service.NewVideoService(ctx).GetPublishList(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Items = &video.Datas{
		Items: list,
		Total: count,
	}

	return
}

// Popular implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Popular(ctx context.Context, req *video.PopularReq) (resp *video.PopularResp, err error) {
	// TODO: Your code here...
	resp = new(video.PopularResp)
	list, err := service.NewVideoService(ctx).GetRankList(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = &video.Data{Items: list}

	return
}

// Search implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Search(ctx context.Context, req *video.SearchReq) (resp *video.SearchResp, err error) {
	// TODO: Your code here...
	resp = new(video.SearchResp)
	list, size, err := service.NewVideoService(ctx).Search(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.Items = &video.Datas{
		Items: pack.BuildVideos(list),
		Total: size,
	}

	return
}

// Info implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Info(ctx context.Context, req *video.InfoReq) (resp *video.InfoResp, err error) {
	// TODO: Your code here...
	resp = new(video.InfoResp)
	info, err := service.NewVideoService(ctx).Info(req)
	if err != nil {
		klog.Error(err)
		resp.Base = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Base = pack.BuildBaseResp(nil)
	resp.Items = &video.Data{Items: pack.BuildVideos(info)}

	return
}
