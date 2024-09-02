package integration

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	accountTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/account"
	accountTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/account"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

var (
	as_token    string
	as_testuser types.User

	accId      int
	as_tsidx   int
	as_instidx int
	as_depidx  int
)

func TestAccountService(t *testing.T) {
	t.Run("Setup", accountServiceSetup)
	t.Run("POST://api/v1/accounts", createAccountTests)
	t.Run("GET://api/v1/accounts", getAccountsTests)
	t.Run("GET://api/v1/accounts/:id", getAccountTests)
	t.Run("PUT://api/v1/accounts/:id", updateAccountTests)
	t.Run("DEL://api/v1/accounts/:id", deleteAccountTests)
	t.Run("Cleanup", accountServiceCleanup)
}

func accountServiceSetup(t *testing.T) {
	as_testuser, as_token = newUserSetup(t)
}

func createAccountTests(t *testing.T) {
	for _, tc := range accountTestCases.GetCreateAccountTests(t, as_testuser.User_id, as_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := as_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			id := accountTester.TestCreateAccount(
				t2,
				tc.Payload,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
			)
			if accId == 0 {
				accId = id
			}
		})
	}
}

func getAccountsTests(t *testing.T) {
	// Get accounts test setup
	t.Run("Setup", func(t2 *testing.T) {
		for i := 0; i < 25; i += 1 {
			accountTester.TestCreateAccount(
				t2,
				account.CreateAccountPayload{
					Name:          fmt.Sprintf("Acc%d", i),
					Tax_shelter:   utils.GetRotatingTaxShelter(&as_tsidx),
					Institution:   utils.GetRotatingInstitution(&as_instidx),
					Is_deprecated: utils.GetRotatingDeprecation(&as_depidx),
				},
				as_token,
				as_testuser.User_id,
				fiber.StatusCreated,
			)
		}
	})

	// Get accounts tests
	for _, tc := range accountTestCases.GetAccountsTestCases(t, as_testuser.User_id, as_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			response, ok := tc.ExpectedResponse.(accountTestCases.GetAccountsExpectedResponse)
			if !ok {
				t.Fatal("invalid GetAccountsExpectedResponse")
			}

			tok := as_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			accountTester.TestGetAccounts(
				t2,
				tc.Route,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
				response,
			)
		})
	}
}

func getAccountTests(t *testing.T) {
	for _, tc := range accountTestCases.GetAccountTests(t, accId, as_testuser.User_id, as_testuser.Email) {
		tok := as_token
		if tc.ReplacementToken != "" {
			tok = tc.ReplacementToken
		}

		t.Run(tc.Title, func(t2 *testing.T) {
			accountTester.TestGetAccount(
				t2,
				tc.ParameterId,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func updateAccountTests(t *testing.T) {
	for _, tc := range accountTestCases.GetUpdateAccountTests(t, accId, as_testuser.User_id, as_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := as_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			accountTester.TestUpdateAccount(
				t2,
				tc.ParameterId,
				tc.Payload,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func deleteAccountTests(t *testing.T) {
	for _, tc := range accountTestCases.DeleteAccountTests(t, accId, as_testuser.User_id, as_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := as_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			accountTester.TestDeleteAccount(
				t2,
				tc.ParameterId,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func accountServiceCleanup(t *testing.T) {
	userTester.TestDeleteUser(t, "", as_token, as_testuser.User_id, fiber.StatusOK)
}
