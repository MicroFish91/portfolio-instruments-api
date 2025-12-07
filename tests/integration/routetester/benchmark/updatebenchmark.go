package benchmark

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBenchmark(t *testing.T, benchmarkId int, payload any, token string, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v2/benchmarks/%d", benchmarkId)

	var updateBenchmarkResponse types.UpdateBenchmarkResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &updateBenchmarkResponse)

	switch expectedStatusCode {
	case 200:
		p := payload.(benchmark.UpdateBenchmarkPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Benchmark{
				Benchmark_id:                benchmarkId,
				Name:                        p.Name,
				Description:                 p.Description,
				Asset_allocation:            p.Asset_allocation,
				Std_dev_pct:                 p.Std_dev_pct,
				Real_return_pct:             p.Real_return_pct,
				Drawdown_yrs:                p.Drawdown_yrs,
				Rec_rebalance_threshold_pct: p.Rec_rebalance_threshold_pct,
				Is_deprecated:               p.Is_deprecated,
				User_id:                     expectedUserId,
				Created_at:                  updateBenchmarkResponse.Data.Benchmark.Created_at,
				Updated_at:                  updateBenchmarkResponse.Data.Benchmark.Updated_at,
			},
			updateBenchmarkResponse.Data.Benchmark,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
