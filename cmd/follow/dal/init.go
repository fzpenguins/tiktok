package dal

import (
	"tiktok/cmd/follow/dal/cache"
	"tiktok/cmd/follow/dal/db"
)

func Init() {
	db.Init()
	cache.InitRedis()
}
