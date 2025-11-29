package benchmark

import (
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateBenchmark(t *testing.T, payload any, token string, expectedUserId int, expectedStatusCode int) int {
	var createBenchmarkResponse types.CreateBenchmarkResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPost, "/api/v1/benchmarks", token, &payload, &createBenchmarkResponse)

	switch expectedStatusCode {
	case 201:
		p := payload.(benchmark.CreateBenchmarkPayload)

		// Account for the fact that rebalance_threshold_pct is optional and has a default value
		if p.Rec_rebalance_threshold_pct == 0 {
			p.Rec_rebalance_threshold_pct = constants.BENCHMARK_REBALANCE_PCT_DEFAULT
		}

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Benchmark{
				Benchmark_id:                createBenchmarkResponse.Data.Benchmark.Benchmark_id,
				Name:                        p.Name,
				Description:                 p.Description,
				Asset_allocation:            p.Asset_allocation,
				Std_dev_pct:                 p.Std_dev_pct,
				Real_return_pct:             p.Real_return_pct,
				Drawdown_yrs:                p.Drawdown_yrs,
				Rec_rebalance_threshold_pct: p.Rec_rebalance_threshold_pct,
				Is_deprecated:               p.Is_deprecated,
				User_id:                     expectedUserId,
				Created_at:                  createBenchmarkResponse.Data.Benchmark.Created_at,
				Updated_at:                  createBenchmarkResponse.Data.Benchmark.Updated_at,
			},
			createBenchmarkResponse.Data.Benchmark,
		)
		return createBenchmarkResponse.Data.Benchmark.Benchmark_id
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return 0
}
