package errno

var (
	Success = NewErrNo(SuccessCode, "Success")

	ServiceError             = NewErrNo(ServiceErrorCode, "service is unable to start successfully")
	ParamError               = NewErrNo(ParamErrorCode, "parameter error")
	AuthorizationFailedError = NewErrNo(AuthorizationFailedErrCode, "authorization failed")
	//User
	UserExistedError = NewErrNo(ParamErrorCode, "user existed")
	UserNotFound     = NewErrNo(ParamErrorCode, "user not found")
	UserMFAInvalid   = NewErrNo(ParamErrorCode, "invalid MFA secret or passcode")

	FileUploadError = NewErrNo(FileUploadErrorCode, "upload meet error")
	NotVideoFile    = NewErrNo(FileUploadError.ErrorCode, "not video file")

	QueryInfoError = NewErrNo(QueryInfoFailed, GetInfoError)
	//SetInfoError = NewErrNo(SetInfoFailed, "set info error")

	LikeExistedError = NewErrNo(ParamErrorCode, "like existed")

	FollowSelfError    = NewErrNo(ParamErrorCode, "can not follow self")
	FollowExistedError = NewErrNo(ParamErrorCode, "follow exist")
)
