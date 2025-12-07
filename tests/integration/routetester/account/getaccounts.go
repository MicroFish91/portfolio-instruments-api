package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/account"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T, route string, token string, expectedUserId int, expectedStatusCode int, expectedResponse account.GetAccountsExpectedResponse) {
	var r string
	if route == "" {
		r = "/api/v2/accounts"
	} else {
		r = route
	}

	var getAccountsResponse types.GetAccountsResponse
	res := utils.SendGetRequest(t, r, token, &getAccountsResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, expectedResponse.Accounts, len(getAccountsResponse.Data.Accounts))
		assert.EqualExportedValues(
			t,
			expectedResponse.Pagination,
			getAccountsResponse.Data.Pagination,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
