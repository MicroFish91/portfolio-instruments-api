package holding

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetHolding(t *testing.T, holdingId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/holdings/%d", holdingId)

	var getHoldingResponse types.GetHoldingResponse
	res := utils.SendGetRequest(t, route, token, &getHoldingResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Holding{
				Holding_id:        holdingId,
				Name:              getHoldingResponse.Data.Holding.Name,
				Ticker:            getHoldingResponse.Data.Holding.Ticker,
				Asset_category:    getHoldingResponse.Data.Holding.Asset_category,
				Expense_ratio_pct: getHoldingResponse.Data.Holding.Expense_ratio_pct,
				Maturation_date:   getHoldingResponse.Data.Holding.Maturation_date,
				Interest_rate_pct: getHoldingResponse.Data.Holding.Interest_rate_pct,
				Is_deprecated:     getHoldingResponse.Data.Holding.Is_deprecated,
				User_id:           expectedUserId,
				Created_at:        getHoldingResponse.Data.Holding.Created_at,
				Updated_at:        getHoldingResponse.Data.Holding.Updated_at,
			},
			getHoldingResponse.Data.Holding,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
