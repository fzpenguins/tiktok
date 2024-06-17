namespace go video



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
    4: required i64 uid,
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

struct BaseResp{
    1: i64 code,
    2: string msg,
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

struct InfoReq{
    1: list<i64> vid,
}

struct InfoResp{
    1: BaseResp base,
    2: Data items,
}

service VideoService{
    FeedResp Feed(1: FeedReq req),
    PublishResp Publish(1: PublishReq req),
    ListResp List(1: ListReq req),
    PopularResp Popular(1: PopularReq req),
    SearchResp Search(1: SearchReq req),
    InfoResp Info(1: InfoReq req),
}