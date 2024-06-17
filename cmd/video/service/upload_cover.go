package service

import minioClient "tiktok/cmd/video/dal/minio"

func (s *VideoService) UploadCover(coverUrl string) (string, error) {
	return minioClient.UploadCover(coverUrl)
}
