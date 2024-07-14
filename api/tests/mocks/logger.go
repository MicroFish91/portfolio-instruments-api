package mocks

import (
	"log/slog"
	"os"
)

func NewMockLogger(logLevel slog.Level) *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return logger
}
