package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ams003010/Copper/api-server/initializers"
	"github.com/ams003010/Copper/api-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/redis/go-redis/v9"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	redisClient *redis.Client
	tracer      trace.Tracer
)

func setupOpenTelemetryTracing() (*sdktrace.TracerProvider, error) {
	// Read tracing endpoint from environment variable
	tracingEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if tracingEndpoint == "" {
		tracingEndpoint = "jaeger:4318" // Default OTLP HTTP endpoint in Docker network
	}

	// Create OTLP HTTP exporter
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(tracingEndpoint),
		otlptracehttp.WithInsecure(), // Use secure connection in production
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	// Create a new TraceProvider with the OTLP exporter
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("copper-api-server"),
			semconv.DeploymentEnvironmentKey.String("development"),
		)),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	// Create a tracer
	tracer = tp.Tracer("copper-api-server")

	return tp, nil
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()

	// Setup OpenTelemetry Tracing
	tp, err := setupOpenTelemetryTracing()
	if err != nil {
		log.Fatalf("Failed to setup OpenTelemetry tracing: %v", err)
	}
	defer tp.Shutdown(context.Background())

	// Initialize Redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	fmt.Println("Connected to Redis!")
}

func main() {
	r := gin.Default()

	// Add tracing middleware
	r.Use(func(c *gin.Context) {
		ctx, span := tracer.Start(c.Request.Context(), "http-request",
			trace.WithAttributes(
				semconv.HTTPMethodKey.String(c.Request.Method),
				semconv.HTTPTargetKey.String(c.Request.URL.Path),
			),
		)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	routes.ImageRegistryRoutes(r, redisClient)

	r.Run()
}
