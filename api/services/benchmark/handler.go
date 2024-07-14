package benchmark

import (
	"log/slog"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type BenchmarkHandlerImpl struct {
	store  types.BenchmarkStore
	logger *slog.Logger
}

func NewBenchmarkHandler(store types.BenchmarkStore, logger *slog.Logger) *BenchmarkHandlerImpl {
	return &BenchmarkHandlerImpl{
		store:  store,
		logger: logger,
	}
}
