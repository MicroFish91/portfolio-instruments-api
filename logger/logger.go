package logger

import (
	"log/slog"
	"os"
)

func NewLogger(logLevel slog.Level) *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return logger
}
