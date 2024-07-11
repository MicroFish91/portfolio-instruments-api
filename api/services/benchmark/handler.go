package benchmark

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type BenchmarkHandlerImpl struct {
	store types.BenchmarkStore
}

func NewBenchmarkHandler(store types.BenchmarkStore) *BenchmarkHandlerImpl {
	return &BenchmarkHandlerImpl{
		store: store,
	}
}
