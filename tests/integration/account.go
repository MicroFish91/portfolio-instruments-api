package integration

import (
	"testing"

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
	t.Run("Setup", setup)

	t.Run("POST://api/v1/accounts", createAccountTestCases)
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
