package tracer

import (
	"fmt"
	"io"
	"tiktok/config"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerconfig "github.com/uber/jaeger-client-go/config"
)

func InitJaegerInKitex(service string) {
	cfg := &jaegerconfig.Configuration{
		Disabled: false,
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.JaegerAddr,
		},
	}

	cfg.ServiceName = service

	tracer, _, err := cfg.NewTracer(
		jaegerconfig.Logger(jaeger.StdLogger),
		jaegerconfig.ZipkinSharedRPCSpan(true),
	)

	if err != nil {
		panic(fmt.Sprintf("cannot init jaeger: %v\n", err))
	}

	opentracing.SetGlobalTracer(tracer)
}

func InitJaegerInHertz(service string) (opentracing.Tracer, io.Closer) {
	cfg := &jaegerconfig.Configuration{
		Disabled: false,
		Sampler: &jaegerconfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerconfig.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: config.JaegerAddr,
		},
	}

	cfg.ServiceName = service

	tracer, closer, err := cfg.NewTracer(
		jaegerconfig.Logger(jaeger.StdLogger),
		jaegerconfig.ZipkinSharedRPCSpan(true),
	)

	if err != nil {
		panic(fmt.Sprintf("cannot init jaeger: %v\n", err))
	}
	return tracer, closer

}
