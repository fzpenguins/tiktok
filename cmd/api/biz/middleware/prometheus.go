package middleware

import (
	"context"
	"tiktok/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		path := string(ctx.Request.Path())
		method := string(ctx.Request.Method())

		timer := prometheus.NewTimer(utils.ResponseDuration.WithLabelValues(method, path))
		ctx.Next(c)
		timer.ObserveDuration()

		statusCode := ctx.Response.StatusCode()
		if statusCode >= 200 && statusCode < 300 {
			utils.RequestCounter.WithLabelValues(method, path).Inc()

		}

	}
}
