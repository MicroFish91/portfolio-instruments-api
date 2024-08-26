package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	holdingTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/holding"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/auth"
	holdingTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/holding"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

var (
	hs_token    string
	hs_testuser types.User
)

func TestHoldingService(t *testing.T) {
	t.Parallel()

	t.Run("Setup", holdingServiceSetup)
	t.Run("POST://api/v1/holdings", createHoldingTests)
}

func holdingServiceSetup(t *testing.T) {
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
		hs_testuser, hs_token = authTester.TestLogin(
			t2,
			auth.LoginPayload{
				Email:    email,
				Password: password,
			},
			fiber.StatusCreated,
		)
	})
}

func createHoldingTests(t *testing.T) {
	for _, tc := range holdingTestCases.GetCreateHoldingTests(t, hs_testuser.User_id, hs_token) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestCreateHolding(
				t2,
				tc.Payload,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}
