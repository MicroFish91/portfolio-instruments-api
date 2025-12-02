package benchmark

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBenchmark(t *testing.T, benchmarkId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/benchmarks/%d", benchmarkId)

	var deleteBenchmarkResponse types.DeleteBenchmarkResponse
	res := utils.SendDeleteRequest(t, route, token, &deleteBenchmarkResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.NotEmpty(t, deleteBenchmarkResponse.Data.Message)
		assert.EqualExportedValues(
			t,
			types.Benchmark{
				Benchmark_id:                benchmarkId,
				Name:                        deleteBenchmarkResponse.Data.Benchmark.Name,
				Description:                 deleteBenchmarkResponse.Data.Benchmark.Description,
				Asset_allocation:            deleteBenchmarkResponse.Data.Benchmark.Asset_allocation,
				Std_dev_pct:                 deleteBenchmarkResponse.Data.Benchmark.Std_dev_pct,
				Real_return_pct:             deleteBenchmarkResponse.Data.Benchmark.Real_return_pct,
				Drawdown_yrs:                deleteBenchmarkResponse.Data.Benchmark.Drawdown_yrs,
				Rec_rebalance_threshold_pct: deleteBenchmarkResponse.Data.Benchmark.Rec_rebalance_threshold_pct,
				Is_deprecated:               deleteBenchmarkResponse.Data.Benchmark.Is_deprecated,
				User_id:                     expectedUserId,
				Created_at:                  deleteBenchmarkResponse.Data.Benchmark.Created_at,
				Updated_at:                  deleteBenchmarkResponse.Data.Benchmark.Updated_at,
			},
			deleteBenchmarkResponse.Data.Benchmark,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
