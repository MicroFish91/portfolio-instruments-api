package holding

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteHolding(t *testing.T, holdingId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/holdings/%d", holdingId)

	var deleteHoldingResponse types.DeleteHoldingResponse
	res := utils.SendDeleteRequest(t, route, token, &deleteHoldingResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.NotEmpty(t, deleteHoldingResponse.Data.Message)
		assert.EqualExportedValues(
			t,
			types.Holding{
				Holding_id:        holdingId,
				Name:              deleteHoldingResponse.Data.Holding.Name,
				Ticker:            deleteHoldingResponse.Data.Holding.Ticker,
				Asset_category:    deleteHoldingResponse.Data.Holding.Asset_category,
				Expense_ratio_pct: deleteHoldingResponse.Data.Holding.Expense_ratio_pct,
				Maturation_date:   deleteHoldingResponse.Data.Holding.Maturation_date,
				Interest_rate_pct: deleteHoldingResponse.Data.Holding.Interest_rate_pct,
				Is_deprecated:     deleteHoldingResponse.Data.Holding.Is_deprecated,
				User_id:           expectedUserId,
				Created_at:        deleteHoldingResponse.Data.Holding.Created_at,
				Updated_at:        deleteHoldingResponse.Data.Holding.Updated_at,
			},
			deleteHoldingResponse.Data.Holding,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
