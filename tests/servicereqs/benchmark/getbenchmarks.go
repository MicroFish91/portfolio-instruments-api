package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetBenchmarks(t *testing.T, route string, token string, expectedUserId int, expectedStatusCode int, expectedResponse benchmark.GetBenchmarksExpectedResponse) {
	var r string
	if route == "" {
		r = "/api/v1/benchmarks"
	} else {
		r = route
	}

	var getBenchmarksResponse types.GetBenchmarksResponse
	res := utils.SendGetRequest(t, r, token, &getBenchmarksResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, expectedResponse.Benchmarks, len(getBenchmarksResponse.Data.Benchmarks))
		assert.EqualExportedValues(
			t,
			expectedResponse.Pagination,
			getBenchmarksResponse.Data.Pagination,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
