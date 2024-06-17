package constants

import "time"

//response的相关设置
const (
	SuccessCode = 10000
	SuccessMsg  = "success"

	FailureCode = -1
	FailureMsg  = "密码错误"
)

//设有accesstoken和refreshtoken的有效时间
const (
	AccessTokenExpireDuration  = time.Hour * 24
	RefreshTokenExpireDuration = time.Hour * 24 * 7
)

//消息类型
const (
	TypePrivateMessage   = "type1"
	TypeGetHistory       = "type2"
	TypeGetUnreadHistory = "type3"
	TypeGroupMessage     = "type4"
	TypeGetGroupHistory  = "type5"
)

//与websocket相关的
const (
	PageSize = 10

	SingleChat               = 0
	GroupChat                = 1
	GetHistoryFromSingleChat = 2
	GetUnreadFromSingleChat  = 3
	GetHistoryFromGroupChat  = 4

	MaxStore = 4 * 1024
)

const (
	APIServiceName         = "api"
	UserServiceName        = "user"
	InteractionServiceName = "interaction"
	FollowServiceName      = "follow"
	ChatServiceName        = "chat"
	VideoServiceName       = "video"
)

//rpc
const (
	MuxConnection  = 10
	RPCTimeout     = 30 * time.Second
	ConnectTimeout = 500 * time.Millisecond
)

const (
	MaxIdleConns    = 10
	MaxConnections  = 1000
	ConnMaxLifetime = 10 * time.Second
)
