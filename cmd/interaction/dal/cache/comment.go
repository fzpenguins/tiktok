package cache

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"sync"
	"tiktok/kitex_gen/interaction"
	"time"
)

func GetVideoCommentCount(ctx context.Context, vid string) (int64, error) {
	return RedisClient.SCard(ctx, VideoCommentCountKey(vid)).Result()
}

func AddVideoCommentCount(ctx context.Context, vid, cid string) error {
	return RedisClient.SAdd(ctx, VideoCommentCountKey(vid), cid).Err()
}

func DecrVideoCommentCount(ctx context.Context, vid, cid string) error {
	return RedisClient.SDiff(ctx, VideoCommentCountKey(vid), cid).Err()
}

func GetChildCommentCount(ctx context.Context, parentId string) (int64, error) {
	return RedisClient.SCard(ctx, ChildCommentCountKey(parentId)).Result()
}

func AddChildCommentCount(ctx context.Context, parentId, cid string) error {
	return RedisClient.SAdd(ctx, ChildCommentCountKey(parentId), cid).Err()
}

func DecrChildCommentCount(ctx context.Context, parentId, cid string) error {
	return RedisClient.SRem(ctx, ChildCommentCountKey(parentId), cid).Err()
}

func AddIntoChildCommentList(ctx context.Context, c *interaction.Comment, parentId string) error {
	m, _ := json.Marshal(c)
	floatCid, _ := strconv.ParseFloat(c.Cid, 64)
	tx := RedisClient.TxPipeline()
	tx.ZAdd(ctx, ChildCommentListKey(parentId), redis.Z{
		Score:  floatCid,
		Member: m,
	})
	tx.Expire(ctx, ChildCommentListKey(parentId), 5*time.Minute)
	_, err := tx.Exec(ctx)
	return err
}

func AddIntoVideoCommentList(ctx context.Context, c *interaction.Comment, vid string) error {
	log.Println(c)
	m, _ := json.Marshal(c)
	floatCid, _ := strconv.ParseFloat(c.Cid, 64)
	tx := RedisClient.TxPipeline()
	tx.ZAdd(ctx, VideoCommentListKey(vid), redis.Z{
		Score:  floatCid,
		Member: m,
	})
	tx.Expire(ctx, VideoCommentListKey(vid), 5*time.Minute)
	_, err := tx.Exec(ctx)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func IsChildCommentListExist(ctx context.Context, parentId string) bool {
	ttl := RedisClient.TTL(ctx, ChildCommentListKey(parentId)).Val()
	return ttl > 0
}

func IsVideoCommentListExist(ctx context.Context, vid string) bool {
	ttl := RedisClient.TTL(ctx, VideoCommentListKey(vid)).Val()
	return ttl > 0
}

func DeleteAllItemInChildCommentList(ctx context.Context, parentId string) error {
	return RedisClient.ZRem(ctx, ChildCommentListKey(parentId), 0, -1).Err()
}

func DeleteAllItemInVideoCommentList(ctx context.Context, vid string) error {
	return RedisClient.ZRem(ctx, VideoCommentListKey(vid), 0, -1).Err()
}

func GetChildCommentList(ctx context.Context, parentId string, pageSize, pageNum int64) []*interaction.Comment {
	members, err := RedisClient.ZRange(ctx, ChildCommentListKey(parentId), 0, -1).Result()
	if err != nil {
		return nil
	}
	log.Println(members, "members")
	resp := make([]*interaction.Comment, len(members))
	var wg sync.WaitGroup
	start := pageNum * pageSize
	for i := start; i < int64(len(members)) && i-start < pageSize; i++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()
			var c *interaction.Comment
			err = json.Unmarshal([]byte(members[index]), &c)
			if err != nil {
				return
			}
			resp[index] = c
		}(i)
	}
	wg.Wait()
	return resp
}

func GetVideoCommentList(ctx context.Context, vid string, pageSize, pageNum int64) []*interaction.Comment {
	members, err := RedisClient.ZRange(ctx, VideoCommentListKey(vid), 0, -1).Result()
	if err != nil {
		return nil
	}
	log.Println("members = ", members)
	resp := make([]*interaction.Comment, len(members))
	var wg sync.WaitGroup
	start := pageNum * pageSize
	for i := start; i < int64(len(members)) && i-start < pageSize; i++ {
		wg.Add(1)
		go func(index int64) {
			defer wg.Done()
			var c *interaction.Comment
			err = json.Unmarshal([]byte(members[index]), &c)
			if err != nil {
				return
			}
			log.Println("c = ", c)
			resp[index] = c
		}(i)
	}
	wg.Wait()
	return resp
}

//
