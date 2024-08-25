package integration

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	accountTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/account"
	accountTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/account"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

var as_token string
var as_testuser types.User

func TestAccountService(t *testing.T) {
	t.Parallel()

	t.Run("Setup", setup)
	t.Run("POST://api/v1/accounts", createAccountTestCases)
	t.Run("GET://api/v1/accounts", getAccountsTestCases)
}

func setup(t *testing.T) {
	email := utils.GetRotatingEmail()
	password := "abcd1234"

	t.Run("Register", func(t2 *testing.T) {
		authTester.TestRegister(
			t2,
			auth.RegisterPayload{
				Email:    email,
				Password: password,
			},
			fiber.StatusCreated,
		)
	})

	t.Run("Login", func(t2 *testing.T) {
		as_testuser, as_token = authTester.TestLogin(
			t2,
			auth.LoginPayload{
				Email:    email,
				Password: password,
			},
			fiber.StatusCreated,
		)
	})
}

func createAccountTestCases(t *testing.T) {
	for _, tc := range accountTestCases.GetCreateAccountTests(t, as_testuser.User_id, as_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := as_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			accountTester.TestCreateAccount(
				t2,
				tc.Payload,
				tok,
				as_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func getAccountsTestCases(t *testing.T) {
	// Get accounts test setup
	t.Run("Setup", func(t2 *testing.T) {
		for i := 0; i < 25; i += 1 {
			accountTester.TestCreateAccount(
				t2,
				account.CreateAccountPayload{
					Name:          fmt.Sprintf("Acc%d", i),
					Tax_shelter:   utils.GetRotatingTaxShelter(),
					Institution:   utils.GetRotatingInstitution(),
					Is_deprecated: utils.GetRotatingDeprecation(),
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
