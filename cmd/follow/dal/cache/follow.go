package cache

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"tiktok/pkg/errno"
)

func IsExistItemInFollowing(ctx context.Context, fromUid string) bool {
	return RedisClient.SCard(ctx, GetToFollower(fromUid)).Val() > 0
}

func IsExistItemInFollower(ctx context.Context, toUid string) bool {
	return RedisClient.SCard(ctx, GetFollower(toUid)).Val() > 0
}

func AddFollow(ctx context.Context, fromUid, toUid string) error {
	if IsFollowingExist(ctx, toUid, fromUid) {
		return errors.New(errno.RepeatOperationError)
	}
	tx := RedisClient.TxPipeline()
	tx.SAdd(ctx, GetFollower(toUid), fromUid)   //对方的follower要更新
	tx.SAdd(ctx, GetToFollower(fromUid), toUid) //自己的关注要更新
	//tx.Expire(ctx, GetFollower(toUid), 5*time.Minute)
	//tx.Expire(ctx, GetToFollower(fromUid), 5*time.Minute)
	_, err := tx.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func IsFollowerExist(ctx context.Context, toUid, fromUid string) bool {
	if !IsFollowerSetExist(ctx, GetFollower(toUid)) {
		return false
	}
	return RedisClient.SIsMember(ctx, GetFollower(toUid), fromUid).Val()
}

func IsFollowerSetExist(ctx context.Context, toUid string) bool {
	t := RedisClient.TTL(ctx, GetFollower(toUid)).Val()
	return t > 0
}

func IsFollowingExist(ctx context.Context, toUid, fromUid string) bool {
	if !IsFollowerSetExist(ctx, GetToFollower(fromUid)) {
		return false
	}
	return RedisClient.SIsMember(ctx, GetToFollower(fromUid), toUid).Val()
}

func IsFollowingSetExist(ctx context.Context, fromUid string) bool {
	t := RedisClient.TTL(ctx, GetToFollower(fromUid)).Val()
	return t > 0
}

func DeleteFollow(ctx context.Context, fromUid, toUid string) error {
	t := RedisClient.TxPipeline()
	if IsFollowingExist(ctx, toUid, fromUid) {
		t.SRem(ctx, GetToFollower(fromUid), toUid)
	}

	if IsFollowerExist(ctx, toUid, fromUid) {
		t.SRem(ctx, GetFollower(toUid), fromUid)
	}

	_, err := t.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func FriendListAction(ctx context.Context, fromUid string) ([]string, error) { //不存在或过期的话，就交给mysql判断
	list, err := RedisClient.SMembers(ctx, GetFollower(fromUid)).Result()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup

	resp := make([]string, len(list))
	for i, s := range list {
		wg.Add(1)
		go func(index int, s string) {
			defer wg.Done()
			if RedisClient.SIsMember(ctx, GetToFollower(fromUid), s).Val() {
				resp[index] = s
			}
		}(i, s)

	}
	wg.Wait()
	return resp, nil
	//return RedisClient.SIsMember(ctx, GetFollower(toUid), fromUid).Val() && RedisClient.SIsMember(ctx, GetFollower(fromUid), toUid).Val()
	//b1 := RedisClient.SIsMember(ctx, GetFollower(toUid), fromUid).Val()
	//b2 := RedisClient.SIsMember(ctx, GetFollower(fromUid), toUid).Val()
	//if b1 && b2 {
	//	return true
	//}
	//
	//b3 := RedisClient.SIsMember(ctx, GetToFollower(fromUid), toUid).Val()
	//if b2 && b3 {
	//	return true
	//}
	//b4 := RedisClient.SIsMember(ctx, GetToFollower(toUid), fromUid).Val()
	//return b1 && b4 || b3 && b4
}

func GetFollowingList(ctx context.Context, pageNum, pageSize int64, fromUid string) ([]string, error) {
	followingList := make([]string, pageSize)
	list, err := RedisClient.SMembers(ctx, GetToFollower(fromUid)).Result()
	if err != nil {
		return nil, err
	}
	start := pageNum * pageSize
	end := start + pageSize - 1

	for i := start; i <= end && i < int64(len(list)); i++ {
		//t, _ := strconv.ParseInt(list[i], 10, 64)
		followingList = append(followingList, list[i])
	}
	return followingList, nil
}

func GetFollowerList(ctx context.Context, pageNum, pageSize int64, toUid string) ([]string, error) {
	followerList := make([]string, pageSize)
	list, err := RedisClient.SMembers(ctx, GetFollower(toUid)).Result()
	if err != nil {
		return nil, err
	}
	start := pageNum * pageSize
	end := start + pageSize - 1

	for i := start; i <= end && i < int64(len(list)); i++ {
		//t, _ := strconv.ParseInt(list[i], 10, 64)
		followerList = append(followerList, list[i])
	}
	return followerList, nil
}
