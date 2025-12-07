package account

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T, accountId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/accounts/%d", accountId)

	var deleteAccountResponse types.DeleteAccountResponse
	res := utils.SendDeleteRequest(t, route, token, &deleteAccountResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.NotEmpty(t, deleteAccountResponse.Data.Message)
		assert.EqualExportedValues(
			t,
			types.Account{
				Account_id:    accountId,
				Name:          deleteAccountResponse.Data.Account.Name,
				Description:   deleteAccountResponse.Data.Account.Description,
				Tax_shelter:   deleteAccountResponse.Data.Account.Tax_shelter,
				Institution:   deleteAccountResponse.Data.Account.Institution,
				Is_deprecated: deleteAccountResponse.Data.Account.Is_deprecated,
				User_id:       expectedUserId,
				Created_at:    deleteAccountResponse.Data.Account.Created_at,
				Updated_at:    deleteAccountResponse.Data.Account.Updated_at,
			},
			deleteAccountResponse.Data.Account,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
