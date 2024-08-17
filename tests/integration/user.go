package integration

import (
	"testing"

	auth "github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	userTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/usercases"
	authTestCase "github.com/MicroFish91/portfolio-instruments-api/tests/testcase/auth"
	userTestCase "github.com/MicroFish91/portfolio-instruments-api/tests/testcase/user"
)

var token string
var user types.User

func TestUser(t *testing.T) {
	t.Parallel()

	t.Run("POST://api/v1/register", registerTestCases)
	t.Run("POST://api/v1/login", loginTestCases)
	t.Run("GET://api/v1/me", getMeTestCases)
}

func registerTestCases(t *testing.T) {
	for _, tc := range userTestCases.RegisterTestCases {
		payload, ok := tc.Payload.(auth.RegisterPayload)
		if !ok {
			t.Fatal("invalid auth register payload")
		}

		t.Run(tc.Title, func(t2 *testing.T) {
			authTestCase.TestRegister(
				t2,
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

		t.Run(tc.Title, func(t2 *testing.T) {
			u, tok := authTestCase.TestLogin(
				t2,
				payload,
				tc.ExpectedStatusCode,
			)
			if token == "" && tok != "" {
				token = tok
				user = u
			}
		})
	}
}

func getMeTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetMeTestCases(t, user.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTestCase.TestGetMe(t2, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}
