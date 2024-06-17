package service

import minioClient "tiktok/cmd/video/dal/minio"

func (s *VideoService) UploadVideo(videoUrl string) (string, error) {
	return minioClient.UploadVideo(videoUrl)
}
