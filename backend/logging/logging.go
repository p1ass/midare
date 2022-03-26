package logging

import (
	"context"
	"fmt"

	"github.com/p1ass/midare/config"
	stackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
)

type ctxLoggerKey struct{}

// New returns struct of zap logger
// If you get logger inside request context, use Extract instead of this function
func New() *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig = stackdriver.EncoderConfig

	logger, _ := cfg.Build()
	return logger
}

// Inject injects logger into context
func Inject(ctx context.Context, logger *zap.Logger) context.Context {
	projectID := config.ReadGoogleCloudProjectID()
	tracer := ExtractTracer(ctx)

	logger = logger.With(
		zap.String("logging.googleapis.com/trace", fmt.Sprintf("projects/%s/traces/%s", projectID, tracer.TraceID)),
		zap.String("logging.googleapis.com/spanId", tracer.SpanID),
	)

	return context.WithValue(ctx, ctxLoggerKey{}, logger)
}

// Extract extracts logger from context
func Extract(ctx context.Context) *zap.Logger {
	return ctx.Value(ctxLoggerKey{}).(*zap.Logger)
}

// Error returns zap field wrapping error
func Error(err error) zap.Field {
	return zap.String("error", fmt.Sprintf("%+v", err))
}
