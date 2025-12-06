package holding

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateHolding(t *testing.T, holdingId int, payload any, token string, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v2/holdings/%d", holdingId)

	var updateAccountResponse types.UpdateHoldingResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &updateAccountResponse)

	switch expectedStatusCode {
	case 200:
		p := payload.(holding.UpdateHoldingPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Holding{
				Holding_id:        holdingId,
				Name:              p.Name,
				Ticker:            p.Ticker,
				Asset_category:    p.Asset_category,
				Expense_ratio_pct: p.Expense_ratio_pct,
				Maturation_date:   p.Maturation_date,
				Interest_rate_pct: p.Interest_rate_pct,
				Is_deprecated:     p.Is_deprecated,
				User_id:           expectedUserId,
				Created_at:        updateAccountResponse.Data.Holding.Created_at,
				Updated_at:        updateAccountResponse.Data.Holding.Updated_at,
			},
			updateAccountResponse.Data.Holding,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
