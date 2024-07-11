package mocks

import "github.com/MicroFish91/portfolio-instruments-api/api/types"

type MockBenchmarkStore struct{}

func NewMockBenchmarkStore() *MockBenchmarkStore {
	return &MockBenchmarkStore{}
}

func (s *MockBenchmarkStore) CreateBenchmark(b *types.Benchmark) error {
	return nil
}
