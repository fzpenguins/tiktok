package main

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"tiktok/cmd/interaction/pack"
	"tiktok/cmd/interaction/service/comment"
	"tiktok/cmd/interaction/service/like"
	interaction "tiktok/kitex_gen/interaction"
	"tiktok/pkg/errno"
)

// InteractionServiceImpl implements the last service interface defined in the IDL.
type InteractionServiceImpl struct{}

// ActionLike implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ActionLike(ctx context.Context, req *interaction.ActionLikeReq) (resp *interaction.ActionLikeResp, err error) {
	// TODO: Your code here...
	resp = new(interaction.ActionLikeResp)
	if req.GetVid() != "" {
		if req.ActionType == "1" {
			err = like.NewLikeService(ctx).LikeVideo(req)
		} else if req.ActionType == "2" {
			err = like.NewLikeService(ctx).DeleteVideoLike(req)
		} else {
			return nil, errno.ParamError
		}
		if err != nil {
			klog.Error(errors.Cause(err))
			resp.Base = pack.BuildBaseResp(err)
			return nil, errors.WithMessage(err, errno.SetInfoError)
		}
	} else if req.GetCid() != "" {
		if req.ActionType == "1" {
			err = like.NewLikeService(ctx).LikeComment(req)
		} else if req.ActionType == "2" {
			err = like.NewLikeService(ctx).DeleteCommentLike(req)
		} else {
			return nil, errno.ParamError
		}
		if err != nil {
			klog.Error(errors.Cause(err))
			resp.Base = pack.BuildBaseResp(err)
			return nil, errors.WithMessage(err, errno.SetInfoError)
		}
	}

	resp.Base = pack.BuildBaseResp(nil)

	return
}

// ListLike implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ListLike(ctx context.Context, req *interaction.ListLikeReq) (resp *interaction.ListLikeResp, err error) {
	// TODO: Your code here...
	resp = new(interaction.ListLikeResp)
	vs, err := like.NewLikeService(ctx).ListLikeVideo(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Data = &interaction.VideosData{Items: vs}
	return
}

// PublishComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) PublishComment(ctx context.Context, req *interaction.PublishCommentReq) (resp *interaction.PublishCommentResp, err error) {
	// TODO: Your code here...
	resp = new(interaction.PublishCommentResp)
	if req.Vid != "" {
		err = comment.NewCommentService(ctx).CreateVideoComment(req)
	} else if req.Cid != "" {
		err = comment.NewCommentService(ctx).CreateChildComment(req)
	} else {
		return nil, errno.ParamError
	}
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.CreateFailed)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// ListComment implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) ListComment(ctx context.Context, req *interaction.ListCommentReq) (resp *interaction.ListCommentResp, err error) {
	// TODO: Your code here...
	resp = new(interaction.ListCommentResp)
	cs := make([]*interaction.Comment, req.GetPageSize())
	if req.GetVid() != "" {
		cs, err = comment.NewCommentService(ctx).VideoCommentList(req)
	} else if req.GetCid() != "" {
		cs, err = comment.NewCommentService(ctx).ChildCommentList(req)
	} else {
		return nil, errno.ParamError
	}

	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}

	resp.Base = pack.BuildBaseResp(nil)
	resp.Items = &interaction.CommentsData{Items: cs}
	return
}

// Delete implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) Delete(ctx context.Context, req *interaction.DeleteReq) (resp *interaction.DeleteResp, err error) {
	// TODO: Your code here...
	resp = new(interaction.DeleteResp)
	err = comment.NewCommentService(ctx).DeleteComment(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.DeleteError)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// GetVideoInfo implements the InteractionServiceImpl interface.
func (s *InteractionServiceImpl) GetVideoInfo(ctx context.Context, req *interaction.GetVideoInfoRequest) (resp *interaction.GetVideoInfoResponse, err error) {
	// TODO: Your code here...
	resp = new(interaction.GetVideoInfoResponse)
	resp, err = like.NewLikeService(ctx).GetVideoInfo(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.GetInfoError)
	}
	resp.Base = pack.BuildBaseResp(nil)

	return
}
