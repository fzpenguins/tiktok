package main

import (
	"context"
	"tiktok/cmd/follow/pack"
	"tiktok/cmd/follow/service"
	follow "tiktok/kitex_gen/follow"
	"tiktok/pkg/errno"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
)

// FollowServiceImpl implements the last service interface defined in the IDL.
type FollowServiceImpl struct{}

// Action implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) Action(ctx context.Context, req *follow.ActionReq) (resp *follow.ActionResp, err error) {
	// TODO: Your code here...

	resp = new(follow.ActionResp)

	if req.ActionType == "0" {
		err = service.NewFollowService(ctx).CreateFollow(req)
	} else if req.ActionType == "1" {
		err = service.NewFollowService(ctx).DeleteFollow(req)
	} else {
		return nil, errno.ParamError
	}
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.FollowOperationError)
	}
	resp.Base = pack.BuildBaseResp(nil)
	return
}

// ListFollowing implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) ListFollowing(ctx context.Context, req *follow.ListFollowingReq) (resp *follow.ListFollowingResp, err error) {
	// TODO: Your code here...
	resp = new(follow.ListFollowingResp)
	resp, err = service.NewFollowService(ctx).FollowList(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}

// ListFollower implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) ListFollower(ctx context.Context, req *follow.ListFollowerReq) (resp *follow.ListFollowerResp, err error) {
	// TODO: Your code here...
	resp = new(follow.ListFollowerResp)
	resp, err = service.NewFollowService(ctx).FollowerList(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return

}

// ListFriend implements the FollowServiceImpl interface.
func (s *FollowServiceImpl) ListFriend(ctx context.Context, req *follow.ListFriendReq) (resp *follow.ListFriendResp, err error) {
	// TODO: Your code here...
	resp = new(follow.ListFriendResp)
	resp, err = service.NewFollowService(ctx).FriendList(req)
	if err != nil {
		klog.Error(errors.Cause(err))
		resp.Base = pack.BuildBaseResp(err)
		return nil, errors.WithMessage(err, errno.QueryFailed)
	}

	resp.Base = pack.BuildBaseResp(nil)
	return
}
