package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"tiktok/cmd/picture/pack"
	"tiktok/cmd/picture/service"
	picture "tiktok/kitex_gen/picture"
)

// PictureServiceImpl implements the last service interface defined in the IDL.
type PictureServiceImpl struct{}

// Insert implements the PictureServiceImpl interface.
func (s *PictureServiceImpl) Insert(ctx context.Context, req *picture.InsertRequest) (resp *picture.InsertResponse, err error) {
	// TODO: Your code here...
	image, err := service.NewPictureService(ctx).Insert(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to insert")
	}

	resp.Image = pack.BuildImage(image)
	resp.Base = pack.BuildBaseResp(nil)
	return
}

// SearchByImage implements the PictureServiceImpl interface.
func (s *PictureServiceImpl) SearchByImage(ctx context.Context, req *picture.SearchByImageRequest) (resp *picture.SearchResponse, err error) {
	// TODO: Your code here...

	images, err := service.NewPictureService(ctx).SearchByImage(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return resp, errors.WithMessage(err, "failed to search")
	}
	resp.Images = pack.BuildImages(images)
	resp.Base = pack.BuildBaseResp(nil)

	return
}
