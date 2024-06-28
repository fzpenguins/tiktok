package errno

const (
	SuccessCode = 10000
	SuccessMsg  = "success"

	FailureCode = -1
	FailureMsg  = "密码错误"

	ServiceErrorCode           = 10001
	ParamErrorCode             = 10002
	AuthorizationFailedErrCode = 10003
	FileUploadErrorCode        = 10009
	QueryInfoFailed            = 10005

	SetInfoFailed  = 10006
	TooManyRequest = 10007
)
