package minio

import (
	"context"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"net/http"
	"path/filepath"
	"tiktok/config"
)

func UploadImage(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	objectName := "picture/" + uuid.Must(uuid.NewRandom()).String() + ext
	resp, err := http.Get(filePath)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, err = MinioClient.PutObject(context.Background(), config.BucketName, objectName, resp.Body, -1, minio.PutObjectOptions{})
	if err != nil {
		return "", err
	}
	return objectName, nil
}
