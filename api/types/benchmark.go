package types

import (
	"time"

	"github.com/gofiber/fiber/v3"
)

type AssetAllocation struct {
	Category AssetCategory `json:"category"`
	Percent  int           `json:"percent"`
}

type Benchmark struct {
	Benchmark_id     int               `json:"benchmark_id"`
	Name             string            `json:"name"`
	Description      string            `json:"description"`
	Asset_allocation []AssetAllocation `json:"asset_allocation"`
	Std_dev_pct      float32           `json:"std_dev_pct"`
	Real_return_pct  float32           `json:"real_return_pct"`
	Drawdown_yrs     int               `json:"drawdown_yrs"`
	Is_deprecated    bool              `json:"is_deprecated"`
	User_id          int               `json:"user_id"`
	Created_at       time.Time         `json:"created_at"`
	Updated_at       time.Time         `json:"updated_at"`
}

type BenchmarkHandler interface {
	CreateBenchmark(fiber.Ctx) error
}

type BenchmarkStore interface {
	CreateBenchmark(*Benchmark) error
}
