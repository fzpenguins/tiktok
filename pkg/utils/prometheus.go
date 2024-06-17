package utils

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "total_http_requests",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint"},
	)
	ResponseDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_response_duration_seconds",
			Help:    "Histogram of response duration for HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)
)

func Init() {
	// prometheus.MustRegister(RequestCounter)
	err := prometheus.Register(RequestCounter)
	if err != nil {
		panic(err)
	}

	err = prometheus.Register(ResponseDuration)
	if err != nil {
		panic(err)
	}

	// prometheus.MustRegister(ResponseDuration)
	log.Println("successfully register prometheus")
}
