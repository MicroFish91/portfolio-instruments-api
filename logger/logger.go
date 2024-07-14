package logger

import (
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: false,
		Level:     slog.LevelDebug, // Todo: Set this based on environment variable
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	return logger
}
