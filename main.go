package main

import (
	"fmt"
	"net/http"

	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"liemlhd.com/tracing-and-monitoring/representative-quotes/tracing"
	"syreclabs.com/go/faker"
)

var (
	serviceName = getEnvWithDefault("JAEGER_SERVICE_NAME", "quotes")
)

func init() {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.EncoderConfig = zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	logger, _ := cfg.Build(
		zap.Fields(zap.String("app_name", serviceName)),
	)

	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

func main() {

	// Init Distributed tracing
	flag := os.Getenv("JAEGER_ENABLED")
	if flag == "true" {
		// 1. init tracer
		tracer, closer := tracing.NewTracer(serviceName)
		if closer != nil {
			defer closer.Close()
		}
		// 2. set the global tracer
		if tracer != nil {
			fmt.Println("Set tracer")
			opentracing.SetGlobalTracer(tracer)
		}
	}

	faker.Seed(int64(100))
	// Echo instance
	e := echo.New()
	e.HideBanner = true
	// Middleware
	if flag == "true" {
		// Tracing
		fmt.Printf("Use middleware with service name: %+v\n", serviceName)
		// 3. use the middleware
		e.Use(tracing.Middleware(serviceName, tracing.NewPathSkipper("/metrics")))
	}
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Skipper: tracing.NewPathSkipper("/metrics"),
	// }))
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	e.GET("/quote/:representative", func(c echo.Context) error {
		representative := c.Param("representative")
		span := opentracing.SpanFromContext(c.Request().Context())
		zap.L().Debug("Receive get quote request.", zap.String("representative", representative), zap.Any("span", span))
		quote := faker.Lorem().Sentence(len(representative))
		return c.String(http.StatusOK, quote)
	})
	// Start server
	e.Logger.Fatal(e.Start(":" + getEnvWithDefault("PORT", "1323")))
}

func getEnvWithDefault(env, deflt string) string {
	val := os.Getenv(env)
	if len(val) == 0 {
		return deflt
	}
	return val
}
