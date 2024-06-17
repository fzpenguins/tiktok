package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"tiktok/config"
)

var MinioClient *minio.Client

func InitMinIoClient() {
	var err error
	MinioClient, err = minio.New(config.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: config.SSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = MinioClient.MakeBucket(context.Background(), config.BucketName, minio.MakeBucketOptions{})
	if err != nil {
		log.Println(err)
	}
}
