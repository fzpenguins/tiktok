package dal

import (
	"tiktok/cmd/api/dal/cache"
	"tiktok/cmd/api/dal/db"
	"tiktok/cmd/api/dal/mq"
)

func Init() {
	cache.InitRedis()
	db.Init()
	mq.LinkRabbitmq()
}
