namespace go interaction

include "video.thrift"

struct BaseResp{
    1: i64 code,
    2: string msg,
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
    4: string uid,
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
    4: required string uid,
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
    3: required string uid,
}

struct GetVideoInfoRequest {
    1: string vid
}


struct VideosData{
    1: list<video.Video> items,
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

struct GetVideoInfoResponse {
    1: BaseResp base,
    2: i64 like_count,
    3: i64 comment_count,
}

struct PublishCommentResp{
    1: BaseResp base,
}

struct DeleteResp{
    1: BaseResp base,
}

service InteractionService{
    ActionLikeResp ActionLike(1: ActionLikeReq req),
    ListLikeResp ListLike(1: ListLikeReq req),
    PublishCommentResp PublishComment(1: PublishCommentReq req),
    ListCommentResp ListComment(1: ListCommentReq req),
    DeleteResp Delete(1: DeleteReq req),
    GetVideoInfoResponse GetVideoInfo(1: GetVideoInfoRequest req),
}