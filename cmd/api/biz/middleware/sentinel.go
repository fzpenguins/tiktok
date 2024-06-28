package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	hertzUtils "github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/opensergo/sentinel/adapter"
	"net/http"
	"tiktok/pkg/errno"
	"tiktok/pkg/utils"
)

func SentinelHandleFunc() app.HandlerFunc {
	return adapter.SentinelServerMiddleware(
		adapter.WithServerResourceExtractor(func(c context.Context, ctx *app.RequestContext) string {
			return "POST/user/register"
		}),

		adapter.WithServerBlockFallback(func(c context.Context, ctx *app.RequestContext) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest,
				hertzUtils.H{
					"err":  "too many request; the quota used up",
					"code": errno.TooManyRequest,
				})
			utils.Logrus.Error(errno.TooManyRequestError)
		}),

	)
}
