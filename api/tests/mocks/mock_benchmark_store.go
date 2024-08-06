package mocks

import (
	"context"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

type MockBenchmarkStore struct{}

func NewMockBenchmarkStore() *MockBenchmarkStore {
	return &MockBenchmarkStore{}
}

func (s *MockBenchmarkStore) CreateBenchmark(ctx context.Context, b types.Benchmark) (types.Benchmark, error) {
	return types.Benchmark{}, nil
}

func (s *MockBenchmarkStore) GetBenchmarks(ctx context.Context, userId int, options types.GetBenchmarksStoreOptions) ([]types.Benchmark, types.PaginationMetadata, error) {
	return []types.Benchmark{
		{
			User_id: userId,
		},
	}, types.PaginationMetadata{}, nil
}

func (s *MockBenchmarkStore) GetBenchmarkById(ctx context.Context, userId, benchmarkId int) (types.Benchmark, error) {
	return types.Benchmark{
		Benchmark_id: benchmarkId,
		User_id:      userId,
	}, nil
}

func (s *MockBenchmarkStore) GetBenchmarkByName(ctx context.Context, name string, userId int) (types.Benchmark, error) {
	return types.Benchmark{
		Name:    name,
		User_id: userId,
	}, nil
}

func (s *MockBenchmarkStore) UpdateBenchmark(ctx context.Context, b types.Benchmark) (types.Benchmark, error) {
	return types.Benchmark{}, nil
}
