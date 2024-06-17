package main

import (
	"context"
	"testing"
	"tiktok/cmd/user/dal"
	"tiktok/cmd/user/service"
	"tiktok/config"
)

var (
	username string
	password string
	uid      string

	access_token, refresh_token string
	userService                 *service.UserService
	avatarUrl                   string
)

func TestMain(m *testing.M) {
	config.ConfigForTest()
	dal.Init()

	userService = service.NewUserService(context.Background())

	username = "penqee"
	password = "123456"
	avatarUrl = "https://th.bing.com/th/id/OIP.VlXsxUWAoGSSgksl1PTANwHaHa?rs=1&pid=ImgDetMain"

	m.Run()
}

func TestMainOrder(t *testing.T) {
	t.Run("register", testRegister)

	t.Run("login", testLogin)

	t.Run("search", testSearch)

	t.Run("upload", testUpload)
}
