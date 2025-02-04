// Code generated by Kitex v0.9.1. DO NOT EDIT.

package userservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "tiktok/kitex_gen/user"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Register": kitex.NewMethodInfo(
		registerHandler,
		newUserServiceRegisterArgs,
		newUserServiceRegisterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Login": kitex.NewMethodInfo(
		loginHandler,
		newUserServiceLoginArgs,
		newUserServiceLoginResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Info": kitex.NewMethodInfo(
		infoHandler,
		newUserServiceInfoArgs,
		newUserServiceInfoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Upload": kitex.NewMethodInfo(
		uploadHandler,
		newUserServiceUploadArgs,
		newUserServiceUploadResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"GetMFA": kitex.NewMethodInfo(
		getMFAHandler,
		newUserServiceGetMFAArgs,
		newUserServiceGetMFAResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"BindMFA": kitex.NewMethodInfo(
		bindMFAHandler,
		newUserServiceBindMFAArgs,
		newUserServiceBindMFAResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	userServiceServiceInfo                = NewServiceInfo()
	userServiceServiceInfoForClient       = NewServiceInfoForClient()
	userServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return userServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return userServiceServiceInfoForClient
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
	serviceName := "UserService"
	handlerType := (*user.UserService)(nil)
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
		"PackageName": "user",
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

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceRegisterArgs)
	realResult := result.(*user.UserServiceRegisterResult)
	success, err := handler.(user.UserService).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceRegisterArgs() interface{} {
	return user.NewUserServiceRegisterArgs()
}

func newUserServiceRegisterResult() interface{} {
	return user.NewUserServiceRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceLoginArgs)
	realResult := result.(*user.UserServiceLoginResult)
	success, err := handler.(user.UserService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceLoginArgs() interface{} {
	return user.NewUserServiceLoginArgs()
}

func newUserServiceLoginResult() interface{} {
	return user.NewUserServiceLoginResult()
}

func infoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceInfoArgs)
	realResult := result.(*user.UserServiceInfoResult)
	success, err := handler.(user.UserService).Info(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceInfoArgs() interface{} {
	return user.NewUserServiceInfoArgs()
}

func newUserServiceInfoResult() interface{} {
	return user.NewUserServiceInfoResult()
}

func uploadHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceUploadArgs)
	realResult := result.(*user.UserServiceUploadResult)
	success, err := handler.(user.UserService).Upload(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceUploadArgs() interface{} {
	return user.NewUserServiceUploadArgs()
}

func newUserServiceUploadResult() interface{} {
	return user.NewUserServiceUploadResult()
}

func getMFAHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceGetMFAArgs)
	realResult := result.(*user.UserServiceGetMFAResult)
	success, err := handler.(user.UserService).GetMFA(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceGetMFAArgs() interface{} {
	return user.NewUserServiceGetMFAArgs()
}

func newUserServiceGetMFAResult() interface{} {
	return user.NewUserServiceGetMFAResult()
}

func bindMFAHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.UserServiceBindMFAArgs)
	realResult := result.(*user.UserServiceBindMFAResult)
	success, err := handler.(user.UserService).BindMFA(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserServiceBindMFAArgs() interface{} {
	return user.NewUserServiceBindMFAArgs()
}

func newUserServiceBindMFAResult() interface{} {
	return user.NewUserServiceBindMFAResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterReq) (r *user.RegisterResp, err error) {
	var _args user.UserServiceRegisterArgs
	_args.Req = req
	var _result user.UserServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *user.LoginReq) (r *user.LoginResp, err error) {
	var _args user.UserServiceLoginArgs
	_args.Req = req
	var _result user.UserServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Info(ctx context.Context, req *user.InfoReq) (r *user.InfoResp, err error) {
	var _args user.UserServiceInfoArgs
	_args.Req = req
	var _result user.UserServiceInfoResult
	if err = p.c.Call(ctx, "Info", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Upload(ctx context.Context, req *user.UploadAvatarUrlReq) (r *user.UploadAvatarUrlResp, err error) {
	var _args user.UserServiceUploadArgs
	_args.Req = req
	var _result user.UserServiceUploadResult
	if err = p.c.Call(ctx, "Upload", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMFA(ctx context.Context, req *user.GetMFAReq) (r *user.GetMFAResp, err error) {
	var _args user.UserServiceGetMFAArgs
	_args.Req = req
	var _result user.UserServiceGetMFAResult
	if err = p.c.Call(ctx, "GetMFA", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) BindMFA(ctx context.Context, req *user.BindMFAReq) (r *user.BindMFAResp, err error) {
	var _args user.UserServiceBindMFAArgs
	_args.Req = req
	var _result user.UserServiceBindMFAResult
	if err = p.c.Call(ctx, "BindMFA", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
