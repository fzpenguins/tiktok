namespace go api

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
}

struct InfoReq{
    1: optional string uid,
}

struct GetMFAReq{

}

struct BindMFAReq{
    1: required string code,
    2: required string secret,
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



struct File{
    1: required string video_url,
    2: required string cover_url,
}



struct Video{
    1: i64 vid,
    2: i64 uid,
    3: string video_url,
    4: string cover_url,
    5: string title,
    6: string description,
    7: i64 visit_count,
    8: i64 like_count,
    9: i64 comment_count,
    10: string created_at,
    11: string updated_at,
    12: string deleted_at,
}

struct PublishReq{
    1: required File data,
    2: required string title,
    3: required string description,
}

struct ListReq{
    1: required string uid,
    2: required i64 page_num,
    3: required i64 page_size,
}

struct PopularReq{
    1: optional i64 page_size,
    2: optional i64 page_num,
}

struct SearchReq{
    1: required string keywords,
    2: required i64 page_size,
    3: required i64 page_num,
    4: optional i64 from_date,
    5: optional i64 to_date,
    6: optional string username,
}

struct FeedReq{
    1: optional i64 time,
}

struct Datas {
    1: list<Video> items,
    2: i64 total,
}

struct Data{
    1: list<Video> items,
}



struct PublishResp{
    1: BaseResp base,
}

struct FeedResp{
    1: BaseResp base,
    2: list<Video> items,
}

struct ListResp{
    1: BaseResp base,
    2: Datas items,
}

struct PopularResp{
    1: BaseResp base,
    2: Data data,
}

struct SearchResp{
    1: BaseResp base,
    2: Datas items,
}


struct Comment{
    1: string uid,
    2: string vid,
    3: string cid,
    4: string parent_id,
    5: i64 like_count,
    6: i64 child_count,
    7: string content,
    8: string created_at,
    9: string updated_at,
    10: string deleted_at,
}

struct ActionLikeReq{
    1: string vid,
    2: string cid,
    3: string action_type,
}

struct ListLikeReq{
    1: string uid,
    2: i64 page_size,
    3: i64 page_num,
}

struct PublishCommentReq{
    1: string vid,
    2: string cid,
    3: required string content,
}

struct ListCommentReq{
        1: string vid,
        2: string cid,
        3: i64 page_size,
        4: i64 page_num,
}

struct DeleteReq{
    1: required string vid,
    2: required string cid,
}

struct VideosData{
    1: list<Video> items,
}

struct CommentsData{
    1: list<Comment> items,
}

struct ListLikeResp{
    1: BaseResp base,
    2: VideosData data,
}

struct ListCommentResp{
    1: BaseResp base,
    2: CommentsData items,
}

struct ActionLikeResp{
    1: BaseResp base,
}

struct PublishCommentResp{
    1: BaseResp base,
}

struct DeleteResp{
    1: BaseResp base,
}




struct UserInfo{
    1: string uid,
    2: string username,
    3: string avatar_url,
}

struct ActionReq{
    1: required string to_uid,
    2: required string action_type,
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















struct Image {
    1: i64 pid
    2: string url
}


struct InsertRequest {
    1: string url
}

struct InsertResponse {
    1: Image image,
    2: BaseResp base;
}

struct SearchByImageRequest {
    1: string url
}


struct SearchResponse {
    1: list<Image> images;
    2: BaseResp base;
}


struct SearchGoodResp{
     1: BaseResp base,
     2: list<Image> images;
}





service UserService{
    RegisterResp Register(1: RegisterReq req) (api.post = "/user/register"),
    LoginResp Login(1: LoginReq req) (api.post = "/user/login"),
    InfoResp Info(1: InfoReq req) (api.get = "/user/info"),
    UploadAvatarUrlResp Upload(1: UploadAvatarUrlReq req) (api.put = "/user/avatar/upload"),
    GetMFAResp GetMFA(1: GetMFAReq req ) (api.get = "/auth/mfa/qrcode"),
    BindMFAResp BindMFA(1: BindMFAReq req) (api.post = "/auth/mfa/bind"),

    InsertResponse Insert(1: InsertRequest req) (api.post="/user/image/insert");
    SearchResponse SearchByImage(1: SearchByImageRequest req) (api.post="/user/image/search");
}

service VideoService{
    FeedResp Feed(1: FeedReq req) (api.get = "/video/feed"),
    PublishResp Publish(1: PublishReq req) (api.post = "/video/publish"),
    ListResp List(1: ListReq req) (api.get = "/video/list"),
    PopularResp Popular(1: PopularReq req) (api.get = "/video/popular"),
    SearchResp Search(1: SearchReq req) (api.post = "/video/search"),
}//PublishReq

service InteractionService{
    ActionLikeResp ActionLike(1: ActionLikeReq req) (api.post = "/like/action"),
    ListLikeResp ListLike(1: ListLikeReq req) (api.get = "/like/list";),
    PublishCommentResp PublishComment(1: PublishCommentReq req) (api.post = "/comment/publish"),
    ListCommentResp ListComment(1: ListCommentReq req) (api.get = "/comment/list"),
    DeleteResp Delete(1: DeleteReq req) (api.delete = "/comment/delete"),
}

service FollowService{
    ActionResp Action(1: ActionReq req) (api.post = "/relation/action"),
    ListFollowingResp ListFollowing(1: ListFollowingReq req) (api.get = "/following/list"),
    ListFollowerResp ListFollower(1: ListFollowerReq req) (api.get = "/follower/list"),
    ListFriendResp ListFriend(1: ListFriendReq req) (api.get = "/friends/list"),

}


