package dal

import (
	"tiktok/cmd/video/dal/cache"
	"tiktok/cmd/video/dal/db"
	minioClient "tiktok/cmd/video/dal/minio"
)

func Init() {
	db.Init()
	cache.InitRedis()
	minioClient.InitMinIoClient()
}
