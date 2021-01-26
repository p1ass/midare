package logging

import (
	"fmt"

	stackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
)

// New returns struct of zap logger
func New() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig = stackdriver.EncoderConfig

	logger, _ := config.Build()
	return logger
}

// Error returns zap field wrapping error
func Error(err error) zap.Field {
	return zap.String("error", fmt.Sprintf("%+v", err))
}
