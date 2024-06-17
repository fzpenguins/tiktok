namespace go follow

struct UserInfo{
    1: string uid,
    2: string username,
    3: string avatar_url,
}

struct ActionReq{
    1: required string to_uid,
    2: required string action_type,
    3: required string uid,
}

struct ListFollowingReq{
    1: required string uid,
    2: optional i64 page_num,
    3: optional i64 page_size,
}

struct ListFollowerReq{
    1: required string uid,
    2: optional i64 page_num,
    3: optional i64 page_size,
}

struct ListFriendReq{
    1: optional i64 page_num,
    2: optional i64 page_size,
    3: required string uid,
}

struct BaseResp{
    1: i64 code,
    2: string msg,
}

struct ActionResp{
    1: BaseResp base,
}

struct UserInfoData{
    1: list<UserInfo> items,
    2: i64 total,
}

struct ListFollowingResp{
    1: BaseResp base,
    2: UserInfoData data,
}

struct ListFollowerResp{
     1: BaseResp base,
     2: UserInfoData data,
}

struct ListFriendResp{
     1: BaseResp base,
     2: UserInfoData data,
}

service FollowService{
    ActionResp Action(1: ActionReq req),
    ListFollowingResp ListFollowing(1: ListFollowingReq req),
    ListFollowerResp ListFollower(1: ListFollowerReq req),
    ListFriendResp ListFriend(1: ListFriendReq req),
}