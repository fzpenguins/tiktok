package cache

import "fmt"

func UserCreatedKey(username string) string {
	return fmt.Sprintf("user %s created at", username)
}

func UserUpdatedKey(username string) string {
	return fmt.Sprintf("user %s updated at", username)
}

func UserMFASecretKey(username string) string {
	return fmt.Sprintf("user %s secret", username)
}

func UserMFACodeUrlKey(username string) string {
	return fmt.Sprintf("user %s codeurl", username)
}
