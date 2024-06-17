namespace go user

struct BaseResp{
    1: i64 code,
    2: string msg,
}

struct User{
    1: i64 uid,
    2: string username,
    3: string avatar_url,
    4: string created_at,
    5: string updated_at,
    6: string deleted_at,

}

struct QRCode{
    1: string secret,
    2: string qrcode,
}

struct RegisterReq{
    1: required string username,
    2: required string password,
}

struct LoginReq{
    1: required string username,
    2: required string password,
    3: optional string code,
}

struct UploadAvatarUrlReq{
    1: required string avatar_url,
    2: required i64 uid,
}

struct InfoReq{
    1: optional string uid,
}

struct GetMFAReq{
     1: required i64 uid,
}

struct BindMFAReq{
    1: required string code,
    2: required string secret,
    3: optional i64 uid,
}

struct Tokens{
    1: string refresh_token,
    2: string access_token,
}

struct RegisterResp{
    1: BaseResp base,
}

struct LoginResp{
    1: BaseResp base,
    2: User data,
    3: Tokens tokens,
}

struct UploadAvatarUrlResp{
    1: BaseResp base,
    2: User data,
}

struct InfoResp{
    1: BaseResp base,
    2: User data,
}

struct GetMFAResp{
    1: BaseResp base,
    2: QRCode data,
}

struct BindMFAResp{
    1: BaseResp base,
}

service UserService{
    RegisterResp Register(1: RegisterReq req),
    LoginResp Login(1: LoginReq req),
    InfoResp Info(1: InfoReq req),
    UploadAvatarUrlResp Upload(1: UploadAvatarUrlReq req),
    GetMFAResp GetMFA(1: GetMFAReq req ),
    BindMFAResp BindMFA(1: BindMFAReq req),
}
