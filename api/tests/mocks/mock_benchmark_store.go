package mocks

import "github.com/MicroFish91/portfolio-instruments-api/api/types"

type MockBenchmarkStore struct{}

func NewMockBenchmarkStore() *MockBenchmarkStore {
	return &MockBenchmarkStore{}
}

func (s *MockBenchmarkStore) CreateBenchmark(b *types.Benchmark) error {
	return nil
}

func (s *MockBenchmarkStore) GetBenchmarks(userId int, options *types.GetBenchmarksStoreOptions) (*[]types.Benchmark, *types.PaginationMetadata, error) {
	return &[]types.Benchmark{
		{
			User_id: userId,
		},
	}, nil, nil
}

func (s *MockBenchmarkStore) GetBenchmarkById(userId, benchmarkId int) (*types.Benchmark, error) {
	return &types.Benchmark{
		Benchmark_id: benchmarkId,
		User_id:      userId,
	}, nil
}
