// Code generated by hertz generator. DO NOT EDIT.

package api

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	api "tiktok/cmd/api/biz/handler/api"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_auth := root.Group("/auth", _authMw()...)
		{
			_mfa := _auth.Group("/mfa", _mfaMw()...)
			_mfa.POST("/bind", append(_bindmfaMw(), api.BindMFA)...)
			_mfa.GET("/qrcode", append(_getmfaMw(), api.GetMFA)...)
		}
	}
	{
		_comment := root.Group("/comment", _commentMw()...)
		_comment.DELETE("/delete", append(_deleteMw(), api.Delete)...)
		_comment.GET("/list", append(_listcommentMw(), api.ListComment)...)
		_comment.POST("/publish", append(_publishcommentMw(), api.PublishComment)...)
	}
	{
		_follower := root.Group("/follower", _followerMw()...)
		_follower.GET("/list", append(_listfollowerMw(), api.ListFollower)...)
	}
	{
		_following := root.Group("/following", _followingMw()...)
		_following.GET("/list", append(_listfollowingMw(), api.ListFollowing)...)
	}
	{
		_friends := root.Group("/friends", _friendsMw()...)
		_friends.GET("/list", append(_listfriendMw(), api.ListFriend)...)
	}
	{
		_like := root.Group("/like", _likeMw()...)
		_like.POST("/action", append(_actionlikeMw(), api.ActionLike)...)
		_like.GET("/list", append(_listlikeMw(), api.ListLike)...)
	}
	{
		_relation := root.Group("/relation", _relationMw()...)
		_relation.POST("/action", append(_actionMw(), api.Action)...)
	}
	{
		_user := root.Group("/user", _userMw()...)
		_user.GET("/info", append(_infoMw(), api.Info)...)
		_user.POST("/login", append(_loginMw(), api.Login)...)
		_user.POST("/register", append(_registerMw(), api.Register)...)
		{
			_avatar := _user.Group("/avatar", _avatarMw()...)
			_avatar.PUT("/upload", append(_uploadMw(), api.Upload)...)
		}
	}
	{
		_video := root.Group("/video", _videoMw()...)
		_video.GET("/feed", append(_feedMw(), api.Feed)...)
		_video.GET("/list", append(_listMw(), api.List)...)
		_video.GET("/popular", append(_popularMw(), api.Popular)...)
		_video.POST("/publish", append(_publishMw(), api.Publish)...)
		_video.POST("/search", append(_searchMw(), api.Search)...)
	}
}
