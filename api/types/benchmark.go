package types

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
)

type AssetAllocationPct struct {
	Category AssetCategory `json:"category"`
	Percent  int           `json:"percent"`
}

type Benchmark struct {
	Benchmark_id     int                  `json:"benchmark_id"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	Asset_allocation []AssetAllocationPct `json:"asset_allocation"`
	Std_dev_pct      float32              `json:"std_dev_pct"`
	Real_return_pct  float32              `json:"real_return_pct"`
	Drawdown_yrs     int                  `json:"drawdown_yrs"`
	Is_deprecated    bool                 `json:"is_deprecated"`
	User_id          int                  `json:"user_id"`
	Created_at       time.Time            `json:"created_at"`
	Updated_at       time.Time            `json:"updated_at"`
}

type BenchmarkHandler interface {
	CreateBenchmark(fiber.Ctx) error
	GetBenchmarks(fiber.Ctx) error
	GetBenchmarkById(fiber.Ctx) error
	UpdateBenchmark(fiber.Ctx) error
	DeleteBenchmark(fiber.Ctx) error
}

type BenchmarkStore interface {
	CreateBenchmark(context.Context, Benchmark) (Benchmark, error)
	GetBenchmarks(ctx context.Context, userId int, options GetBenchmarksStoreOptions) ([]Benchmark, PaginationMetadata, error)
	GetBenchmarkById(ctx context.Context, userId, benchmarkId int) (Benchmark, error)
	GetBenchmarkByName(ctx context.Context, name string, userId int) (Benchmark, error)
	UpdateBenchmark(context.Context, Benchmark) (Benchmark, error)
	DeleteBenchmark(ctx context.Context, userId, benchmarkId int) (Benchmark, error)
}

type GetBenchmarksStoreOptions struct {
	Benchmark_ids []int
	Name          string
	Is_deprecated string
	Current_page  int
	Page_size     int
}

// ---- Benchmark Response Types ----
type CreateBenchmarkResponse struct {
	Data struct {
		Benchmark Benchmark `json:"benchmark"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetBenchmarksResponse struct {
	Data struct {
		Benchmarks []Benchmark        `json:"benchmarks"`
		Pagination PaginationMetadata `json:"pagination"`
	} `json:"data"`
	Error string `json:"error"`
}

type GetBenchmarkResponse struct {
	Data struct {
		Benchmark Benchmark `json:"benchmark"`
	} `json:"data"`
	Error string `json:"error"`
}

type UpdateBenchmarkResponse struct {
	Data struct {
		Benchmark Benchmark `json:"benchmark"`
	} `json:"data"`
	Error string `json:"error"`
}
