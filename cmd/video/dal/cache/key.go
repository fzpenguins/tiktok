package cache

import (
	"fmt"
	"strconv"
)

func GetVideoLikeCountKey(vid int64) string {
	return fmt.Sprintf("like_count/video:%s", strconv.Itoa(int(vid)))
}

func GetVideoLikeFromUser(uid int64) string {
	return fmt.Sprintf("%s/like_count/video", strconv.Itoa(int(uid))) //传入点赞的人的uid
}

func GetVideoVisitCountKey(vid int64) string {
	return fmt.Sprintf("visit_count/video:%s", strconv.Itoa(int(vid)))
}

func VideoInfoKey(vid string) string {
	return fmt.Sprintf("video:info:%s", vid)
}

// VideoViewsKey 视频观看数
func VideoViewsKey(vid string) string {
	return fmt.Sprintf("video:views:%s", vid)
}
