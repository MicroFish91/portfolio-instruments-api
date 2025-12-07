package account

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccount(t *testing.T, accountId int, payload any, token string, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v2/accounts/%d", accountId)

	var updateAccountResponse types.UpdateAccountResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &updateAccountResponse)

	switch expectedStatusCode {
	case 200:
		p := payload.(account.UpdateAccountPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Account{
				Account_id:    accountId,
				Name:          p.Name,
				Description:   p.Description,
				Tax_shelter:   p.Tax_shelter,
				Institution:   p.Institution,
				Is_deprecated: p.Is_deprecated,
				User_id:       expectedUserId,
				Created_at:    updateAccountResponse.Data.Account.Created_at,
				Updated_at:    updateAccountResponse.Data.Account.Updated_at,
			},
			updateAccountResponse.Data.Account,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
