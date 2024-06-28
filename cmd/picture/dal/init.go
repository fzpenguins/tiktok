package dal

import (
	"tiktok/cmd/picture/dal/db"
	"tiktok/cmd/picture/dal/milvus"
	"tiktok/cmd/picture/dal/minio"
)

func Init() {
	db.Init()
	milvus.Init()
	minio.InitMinIoClient()
}
