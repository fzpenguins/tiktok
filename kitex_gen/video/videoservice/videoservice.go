// Code generated by Kitex v0.9.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	video "tiktok/kitex_gen/video"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Feed": kitex.NewMethodInfo(
		feedHandler,
		newVideoServiceFeedArgs,
		newVideoServiceFeedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Publish": kitex.NewMethodInfo(
		publishHandler,
		newVideoServicePublishArgs,
		newVideoServicePublishResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"List": kitex.NewMethodInfo(
		listHandler,
		newVideoServiceListArgs,
		newVideoServiceListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Popular": kitex.NewMethodInfo(
		popularHandler,
		newVideoServicePopularArgs,
		newVideoServicePopularResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Search": kitex.NewMethodInfo(
		searchHandler,
		newVideoServiceSearchArgs,
		newVideoServiceSearchResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Info": kitex.NewMethodInfo(
		infoHandler,
		newVideoServiceInfoArgs,
		newVideoServiceInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	videoServiceServiceInfo                = NewServiceInfo()
	videoServiceServiceInfoForClient       = NewServiceInfoForClient()
	videoServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return videoServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "video",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceFeedArgs)
	realResult := result.(*video.VideoServiceFeedResult)
	success, err := handler.(video.VideoService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceFeedArgs() interface{} {
	return video.NewVideoServiceFeedArgs()
}

func newVideoServiceFeedResult() interface{} {
	return video.NewVideoServiceFeedResult()
}

func publishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePublishArgs)
	realResult := result.(*video.VideoServicePublishResult)
	success, err := handler.(video.VideoService).Publish(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePublishArgs() interface{} {
	return video.NewVideoServicePublishArgs()
}

func newVideoServicePublishResult() interface{} {
	return video.NewVideoServicePublishResult()
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceListArgs)
	realResult := result.(*video.VideoServiceListResult)
	success, err := handler.(video.VideoService).List(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceListArgs() interface{} {
	return video.NewVideoServiceListArgs()
}

func newVideoServiceListResult() interface{} {
	return video.NewVideoServiceListResult()
}

func popularHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServicePopularArgs)
	realResult := result.(*video.VideoServicePopularResult)
	success, err := handler.(video.VideoService).Popular(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServicePopularArgs() interface{} {
	return video.NewVideoServicePopularArgs()
}

func newVideoServicePopularResult() interface{} {
	return video.NewVideoServicePopularResult()
}

func searchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceSearchArgs)
	realResult := result.(*video.VideoServiceSearchResult)
	success, err := handler.(video.VideoService).Search(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceSearchArgs() interface{} {
	return video.NewVideoServiceSearchArgs()
}

func newVideoServiceSearchResult() interface{} {
	return video.NewVideoServiceSearchResult()
}

func infoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*video.VideoServiceInfoArgs)
	realResult := result.(*video.VideoServiceInfoResult)
	success, err := handler.(video.VideoService).Info(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newVideoServiceInfoArgs() interface{} {
	return video.NewVideoServiceInfoArgs()
}

func newVideoServiceInfoResult() interface{} {
	return video.NewVideoServiceInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *video.FeedReq) (r *video.FeedResp, err error) {
	var _args video.VideoServiceFeedArgs
	_args.Req = req
	var _result video.VideoServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Publish(ctx context.Context, req *video.PublishReq) (r *video.PublishResp, err error) {
	var _args video.VideoServicePublishArgs
	_args.Req = req
	var _result video.VideoServicePublishResult
	if err = p.c.Call(ctx, "Publish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, req *video.ListReq) (r *video.ListResp, err error) {
	var _args video.VideoServiceListArgs
	_args.Req = req
	var _result video.VideoServiceListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Popular(ctx context.Context, req *video.PopularReq) (r *video.PopularResp, err error) {
	var _args video.VideoServicePopularArgs
	_args.Req = req
	var _result video.VideoServicePopularResult
	if err = p.c.Call(ctx, "Popular", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Search(ctx context.Context, req *video.SearchReq) (r *video.SearchResp, err error) {
	var _args video.VideoServiceSearchArgs
	_args.Req = req
	var _result video.VideoServiceSearchResult
	if err = p.c.Call(ctx, "Search", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Info(ctx context.Context, req *video.InfoReq) (r *video.InfoResp, err error) {
	var _args video.VideoServiceInfoArgs
	_args.Req = req
	var _result video.VideoServiceInfoResult
	if err = p.c.Call(ctx, "Info", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
