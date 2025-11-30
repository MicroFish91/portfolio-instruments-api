package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type GetBenchmarksExpectedResponse struct {
	Benchmarks int
	Pagination types.PaginationMetadata
}

func TestGetBenchmarks(t *testing.T, route string, token string, expectedUserId int, expectedStatusCode int, expectedResponse GetBenchmarksExpectedResponse) {
	var r string
	if route == "" {
		r = "/api/v2/benchmarks"
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
