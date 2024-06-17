package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"tiktok/config"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
		//Username:    "127.0.0.1:6379",
		Password:    config.RedisPassword,
		DB:          config.RedisDB,
		ReadTimeout: 5 * time.Second,
		DialTimeout: 5 * time.Second,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
