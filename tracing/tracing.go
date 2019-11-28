package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
)

// NewTracer create new tracer
func NewTracer(defaultServiceName string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Println("cannot parse jaeger env vars: %v\n", err.Error())
		//os.Exit(1)
		return nil, nil
	}
	if len(cfg.ServiceName) == 0 {
		cfg.ServiceName = defaultServiceName
	}
	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer, err := cfg.New(defaultServiceName, config.ZipkinSharedRPCSpan(true),
		config.Injector(opentracing.HTTPHeaders, zipkinPropagator),
		config.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
	)
	if err != nil {
		fmt.Printf("cannot initialize jaeger tracer: %+v\n", err.Error())
		//os.Exit(1)
		return nil, nil
	}
	return tracer, closer
}
