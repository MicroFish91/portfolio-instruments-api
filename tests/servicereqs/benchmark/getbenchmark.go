package benchmark

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetBenchmark(t *testing.T, benchmarkId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/benchmarks/%d", benchmarkId)

	var getBenchmarkResponse types.GetBenchmarkResponse
	res := utils.SendGetRequest(t, route, token, &getBenchmarkResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Benchmark{
				Benchmark_id:     benchmarkId,
				Name:             getBenchmarkResponse.Data.Benchmark.Name,
				Description:      getBenchmarkResponse.Data.Benchmark.Description,
				Asset_allocation: getBenchmarkResponse.Data.Benchmark.Asset_allocation,
				Std_dev_pct:      getBenchmarkResponse.Data.Benchmark.Std_dev_pct,
				Real_return_pct:  getBenchmarkResponse.Data.Benchmark.Real_return_pct,
				Drawdown_yrs:     getBenchmarkResponse.Data.Benchmark.Drawdown_yrs,
				Is_deprecated:    getBenchmarkResponse.Data.Benchmark.Is_deprecated,
				User_id:          expectedUserId,
				Created_at:       getBenchmarkResponse.Data.Benchmark.Created_at,
				Updated_at:       getBenchmarkResponse.Data.Benchmark.Updated_at,
			},
			getBenchmarkResponse.Data.Benchmark,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
