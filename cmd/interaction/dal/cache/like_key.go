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

func GetCommentLikeCountKey(cid int64) string {
	return fmt.Sprintf("like_count/comment:%s", strconv.Itoa(int(cid)))
}

func GetLikeFromUser(uid int64) string {
	return fmt.Sprintf("%s/like_count/comment", strconv.Itoa(int(uid))) //传入点赞的人的uid
}
