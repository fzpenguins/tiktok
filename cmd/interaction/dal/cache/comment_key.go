package cache

import "fmt"

func VideoCommentCountKey(vid string) string {
	return fmt.Sprintf("video_comment_count:%s", vid)
}

func ChildCommentCountKey(cid string) string {
	return fmt.Sprintf("child_comment_count:%s", cid)
}

func ChildCommentListKey(parentId string) string {
	return fmt.Sprintf("child_comment_list:%s", parentId)
}

func VideoCommentListKey(vid string) string {
	return fmt.Sprintf("video_comment_list:%s", vid)
}
