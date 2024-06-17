package dal

import (
	"tiktok/cmd/user/dal/cache"
	"tiktok/cmd/user/dal/db"
	minioclient "tiktok/cmd/user/dal/minio"
)

func Init() {
	db.Init()
	cache.InitRedis()
	minioclient.InitMinIoClient()
}
