package cache

import (
	"context"
	"strconv"
)

func SetCreatedTime(ctx context.Context, username, createTime string) error {
	return RedisClient.HSet(ctx, UserCreatedKey(username), createTime).Err()
}

func SetUpdatedTime(ctx context.Context, username, updatedTime string) error {
	return RedisClient.HSet(ctx, UserCreatedKey(username), updatedTime).Err()
}

func SetUserMFASecret(ctx context.Context, username, secret string) error {
	return RedisClient.Set(ctx, UserMFASecretKey(username), secret, 0).Err()
}

func SetUserMFACodeUrl(ctx context.Context, username, codeUrl string) error {
	return RedisClient.Set(ctx, UserMFACodeUrlKey(username), codeUrl, 0).Err()
}

func SetUsernameHashUid(ctx context.Context, username string, uid int64) error {
	return RedisClient.Set(ctx, username, strconv.FormatInt(uid, 10), 0).Err()
}
