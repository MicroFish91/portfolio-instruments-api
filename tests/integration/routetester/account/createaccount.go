package account

import (
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T, payload any, token string, expectedUserId int, expectedStatusCode int) int {
	var createAccountResponse types.CreateAccountResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPost, "/api/v2/accounts", token, &payload, &createAccountResponse)

	switch expectedStatusCode {
	case 201:
		p := payload.(account.CreateAccountPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Account{
				Account_id:    createAccountResponse.Data.Account.Account_id,
				Name:          p.Name,
				Description:   p.Description,
				Tax_shelter:   p.Tax_shelter,
				Institution:   p.Institution,
				Is_deprecated: p.Is_deprecated,
				User_id:       expectedUserId,
				Created_at:    createAccountResponse.Data.Account.Created_at,
				Updated_at:    createAccountResponse.Data.Account.Updated_at,
			},
			createAccountResponse.Data.Account,
		)
		return createAccountResponse.Data.Account.Account_id
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return 0
}
