package cache

import "fmt"

func GetFollower(toUid string) string {
	return fmt.Sprintf("%s follower:", toUid)
}

func GetToFollower(fromUid string) string {
	return fmt.Sprintf("%s following:", fromUid)
}
