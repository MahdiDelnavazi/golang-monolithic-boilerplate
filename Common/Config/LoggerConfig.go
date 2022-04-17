package Config

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger Create new zap logger instance for handling logging in info level and above

var Logger *zap.SugaredLogger

func NewLogger(service string, Host string) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	config.InitialFields = map[string]interface{}{
		"service": service,
	}
	log, err := config.Build()

	if err != nil {
		fmt.Errorf("error at start %w", err)
	}
	Logger = log.Sugar()

	defer Logger.Sync()
	defer Logger.Infow("shutdown", "status", "here", "host", Host)
}
