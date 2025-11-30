package holding

import (
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateHolding(t *testing.T, payload any, token string, expectedUserId int, expectedStatusCode int) int {
	var createHoldingResponse types.CreateHoldingResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPost, "/api/v2/holdings", token, &payload, &createHoldingResponse)

	switch expectedStatusCode {
	case 201:
		p := payload.(holding.CreateHoldingPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Holding{
				Holding_id:        createHoldingResponse.Data.Holding.Holding_id,
				Name:              p.Name,
				Ticker:            p.Ticker,
				Asset_category:    p.Asset_category,
				Expense_ratio_pct: p.Expense_ratio_pct,
				Maturation_date:   p.Maturation_date,
				Interest_rate_pct: p.Interest_rate_pct,
				Is_deprecated:     p.Is_deprecated,
				User_id:           expectedUserId,
				Created_at:        createHoldingResponse.Data.Holding.Created_at,
				Updated_at:        createHoldingResponse.Data.Holding.Updated_at,
			},
			createHoldingResponse.Data.Holding,
		)
		return createHoldingResponse.Data.Holding.Holding_id
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return 0
}
