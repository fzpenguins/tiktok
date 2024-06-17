package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
)

//func IsLikeCountTimeOut(ctx context.Context, vid int64) bool {
//	ttl := RedisClient.TTL(ctx, GetVideoLikeCountKey(vid)).Val()
//	return ttl <= 0
//}

//func IsUserLikeTimeOut(ctx context.Context, uid int64) bool {
//	return RedisClient.TTL(ctx, GetVideoLikeFromUser(uid)).Val() <= 0
//}

func LikeCount(ctx context.Context, vid int64) int64 {
	countStr, _ := RedisClient.Get(ctx, GetVideoLikeCountKey(vid)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

func AddLikeCount(ctx context.Context, uid int64, vid int64) {
	if RedisClient.ZScore(ctx, GetVideoLikeFromUser(uid), strconv.FormatInt(vid, 10)).Val() != 1 {
		RedisClient.Incr(ctx, GetVideoLikeCountKey(vid))
		RedisClient.ZIncrBy(ctx, "like_count", 1, strconv.Itoa(int(vid)))
	}
	log.Println(1)
	RedisClient.ZAdd(ctx, GetVideoLikeFromUser(uid), redis.Z{Member: strconv.FormatInt(vid, 10), Score: 1})
	log.Println(2)
	//RedisClient.Expire(ctx, GetVideoLikeFromUser(uid), 5*time.Minute)
	//RedisClient.Expire(ctx, GetVideoLikeCountKey(vid), 5*time.Minute)
}

func DecrLikeCount(ctx context.Context, uid int64, vid int64) {
	if RedisClient.ZScore(ctx, GetVideoLikeFromUser(uid), strconv.FormatInt(vid, 10)).Val() != 0 {
		RedisClient.Decr(ctx, GetVideoLikeCountKey(vid))
		RedisClient.ZIncrBy(ctx, "like_count", -1, strconv.Itoa(int(vid))) //这一步可能没用？
	}
	RedisClient.ZAdd(ctx, GetVideoLikeFromUser(uid), redis.Z{Member: strconv.FormatInt(vid, 10), Score: 0})

	//RedisClient.Expire(ctx, GetVideoLikeFromUser(uid), 5*time.Minute)
	//RedisClient.Expire(ctx, GetVideoLikeCountKey(vid), 5*time.Minute)
	//RedisClient.Expire()
}

//func SetLikeCount(ctx context.Context, vid, likeCount int64) error {
//	return RedisClient.Set(ctx, GetVideoLikeCountKey(vid), likeCount, 5*time.Minute).Err()
//}

//func SetUserLike(ctx context.Context, uid, vid int64, isLike string) error {
//	if isLike == "0" {
//		return nil
//	}
//	err := RedisClient.ZAdd(ctx, GetVideoLikeFromUser(uid), redis.Z{Member: vid, Score: 1}).Err()
//	RedisClient.Expire(ctx, GetVideoLikeFromUser(uid), 5*time.Minute)
//	return err
//}

//comment

//func IsCommentLikeCountTimeOut(ctx context.Context, cid int64) bool {
//	ttl := RedisClient.TTL(ctx, GetVideoLikeCountKey(cid)).Val()
//	return ttl <= 0
//}

//func IsUserLikeCommentTimeOut(ctx context.Context, uid int64) bool {
//	return RedisClient.TTL(ctx, GetLikeFromUser(uid)).Val() <= 0
//}

func AddCommentLikeCount(ctx context.Context, uid, cid int64) {

	if RedisClient.ZScore(ctx, GetLikeFromUser(uid), strconv.FormatInt(cid, 10)).Val() != 1 {

		RedisClient.Incr(ctx, GetVideoLikeCountKey(cid))
		RedisClient.ZIncrBy(ctx, "comment:like_count", 1, strconv.Itoa(int(cid)))
	}
	log.Println(1)
	RedisClient.ZAdd(ctx, GetLikeFromUser(uid), redis.Z{Member: strconv.FormatInt(cid, 10), Score: 1}) //comment.Cid, 1) //自己的页面
	log.Println(2)
	//RedisClient.Expire(ctx, GetLikeFromUser(uid), 5*time.Minute)
	//RedisClient.Expire(ctx, GetVideoLikeCountKey(cid), 5*time.Minute)
}

func CommentLikeCount(ctx context.Context, cid int64) int64 {
	countStr, _ := RedisClient.Get(ctx, GetCommentLikeCountKey(cid)).Result()
	count, _ := strconv.ParseInt(countStr, 10, 64)
	return count
}

func DecrCommentLikeCount(ctx context.Context, uid, cid int64) {
	if RedisClient.ZScore(ctx, GetLikeFromUser(uid), strconv.FormatInt(cid, 10)).Val() != 0 {
		//还要有一个所有人可见的页面
		RedisClient.Decr(ctx, GetVideoLikeCountKey(cid))
		RedisClient.ZIncrBy(ctx, "comment:like_count", -1, strconv.Itoa(int(cid)))
	}
	RedisClient.ZAdd(ctx, GetLikeFromUser(uid), redis.Z{Member: strconv.FormatInt(cid, 10), Score: 0})

	//RedisClient.Expire(ctx, GetLikeFromUser(uid), 5*time.Minute)
	//RedisClient.Expire(ctx, GetVideoLikeCountKey(cid), 5*time.Minute)
}
