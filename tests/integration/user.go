package integration

import (
	"testing"

	auth "github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	userTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/usercases"
	authTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/testcases/auth"
)

func TestUser(t *testing.T) {
	t.Parallel()

	t.Run("POST://api/v1/register", registerTestCases)
	t.Run("POST://api/v1/login", loginTestCases)
}

func registerTestCases(t *testing.T) {
	for _, tc := range userTestCases.RegisterTestCases {
		payload, ok := tc.Payload.(auth.RegisterPayload)
		if !ok {
			t.Fatal("invalid auth register payload")
		}

		t.Run(tc.Title, func(t *testing.T) {
			authTestCases.TestRegister(
				t,
				payload,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func loginTestCases(t *testing.T) {
	for _, tc := range userTestCases.LoginTestCases {
		payload, ok := tc.Payload.(auth.LoginPayload)
		if !ok {
			t.Fatal("invalid auth login payload")
		}

		t.Run(tc.Title, func(t *testing.T) {
			authTestCases.TestLogin(
				t,
				payload,
				tc.ExpectedStatusCode,
			)
		})
	}
}
