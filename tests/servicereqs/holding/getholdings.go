package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type GetHoldingsExpectedResponse struct {
	Holdings   int
	Pagination types.PaginationMetadata
}

func TestGetHoldings(t *testing.T, route string, token string, expectedUserId int, expectedStatusCode int, expectedResponse GetHoldingsExpectedResponse) {
	var r string
	if route == "" {
		r = "/api/v1/holdings"
	} else {
		r = route
	}

	var getHoldingsResponse types.GetHoldingsResponse
	res := utils.SendGetRequest(t, r, token, &getHoldingsResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, expectedResponse.Holdings, len(getHoldingsResponse.Data.Holdings))
		assert.EqualExportedValues(
			t,
			expectedResponse.Pagination,
			getHoldingsResponse.Data.Pagination,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
