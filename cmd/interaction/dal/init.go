package dal

import (
	"tiktok/cmd/interaction/dal/cache"
	"tiktok/cmd/interaction/dal/db"
)

func Init() {
	db.Init()
	cache.InitRedis()
}
