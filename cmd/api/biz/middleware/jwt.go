package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
	"net/http"
	"tiktok/pkg/errno"
	"tiktok/pkg/utils"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var code int
		code = http.StatusOK
		accessToken := c.GetHeader("access_token")
		refreshToken := c.GetHeader("refresh_token")
		if string(accessToken) == "" {

			code = http.StatusBadRequest
			c.JSON(http.StatusOK, map[string]interface{}{
				"status": code,
				"data":   "Token can not be empty",
				"error":  errno.AuthorizationFailedError,
			})
			c.Abort()
			return
		}
		newAccessToken, newRefreshToken, err := utils.ParseRefreshToken(string(accessToken), string(refreshToken))
		if err != nil {
			code = http.StatusBadRequest
		}
		if code != http.StatusOK {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status": code,
				"data":   errno.ParseFailed,
				"error":  errors.WithMessage(errno.AuthorizationFailedError, err.Error()), //err.Error(),
			})
			c.Abort()
			return
		}
		c.Header("access_token", newAccessToken)
		c.Header("refresh_token", newRefreshToken)

		//claims, err := utils.ParseToken(newAccessToken)
		//if err != nil {
		//	c.JSON(http.StatusOK, map[string]interface{}{
		//		"status": code,
		//		"data":   errno.ParseFailed,
		//		"error":  errors.WithMessage(errno.AuthorizationFailedError, err.Error()),
		//	})
		//	c.Abort()
		//	return
		//}
		//
		//ctx = pack.NewContext(ctx, &pack.UserInfo{ID: claims.Uid,
		//	UserName: claims.UserName,
		//})

		c.Next(ctx)
	}
}
