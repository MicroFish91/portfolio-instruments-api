package account

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T, accountId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/accounts/%d", accountId)

	var getAccountResponse types.GetAccountResponse
	res := utils.SendGetRequest(t, route, token, &getAccountResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Account{
				Account_id:    accountId,
				Name:          getAccountResponse.Data.Account.Name,
				Description:   getAccountResponse.Data.Account.Description,
				Tax_shelter:   getAccountResponse.Data.Account.Tax_shelter,
				Institution:   getAccountResponse.Data.Account.Institution,
				Is_deprecated: getAccountResponse.Data.Account.Is_deprecated,
				User_id:       expectedUserId,
				Created_at:    getAccountResponse.Data.Account.Created_at,
				Updated_at:    getAccountResponse.Data.Account.Updated_at,
			},
			getAccountResponse.Data.Account,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
