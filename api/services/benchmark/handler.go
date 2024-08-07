package benchmark

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type BenchmarkHandlerImpl struct {
	userStore      types.UserStore
	benchmarkStore types.BenchmarkStore
	logger         *slog.Logger
}

func NewBenchmarkHandler(userStore types.UserStore, benchmarkStore types.BenchmarkStore, logger *slog.Logger) *BenchmarkHandlerImpl {
	return &BenchmarkHandlerImpl{
		userStore:      userStore,
		benchmarkStore: benchmarkStore,
		logger:         logger,
	}
}
