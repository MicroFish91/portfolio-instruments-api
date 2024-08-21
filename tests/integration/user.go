package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	userTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/user"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/auth"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
)

var us_token string
var us_testuser types.User

var email string = utils.GetRotatingEmail()
var password string = "abcd1234"

func TestUserService(t *testing.T) {
	t.Run("POST://api/v1/register", registerTestCases)
	t.Run("POST://api/v1/login", loginTestCases)
	t.Run("GET://api/v1/me", getMeTestCases)
	t.Run("GET://api/v1/users/:id", getUserByIdTestCases)
	t.Run("GET://api/v1/users/:id/settings", getSettingsTestCases)
	t.Run("PUT://api/v1/users/:id/settings", updateSettingsTestCases)
	t.Run("DEL://api/v1/users/:id", deleteUserTestCases)
}

func registerTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetRegisterTestCases(email, password) {
		payload, ok := tc.Payload.(auth.RegisterPayload)
		if !ok {
			t.Fatal("invalid auth register payload")
		}

		t.Run(tc.Title, func(t2 *testing.T) {
			authTester.TestRegister(
				t2,
				payload,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func loginTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetLoginTestCases(email, password) {
		payload, ok := tc.Payload.(auth.LoginPayload)
		if !ok {
			t.Fatal("invalid auth login payload")
		}

		t.Run(tc.Title, func(t2 *testing.T) {
			u, tok := authTester.TestLogin(
				t2,
				payload,
				tc.ExpectedStatusCode,
			)
			if us_token == "" && tok != "" {
				us_token = tok
				us_testuser = u
			}
		})
	}
}

func getMeTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetMeTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetMe(t2, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func getUserByIdTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetUserByIdTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetUserById(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func getSettingsTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetSettingsTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetSettings(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func updateSettingsTestCases(t *testing.T) {
	for _, tc := range userTestCases.UpdateSettingsTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestUpdateSettings(t2, tc.Route, tok, tc.ParameterId, tc.Payload, tc.ExpectedStatusCode)
		})
	}
}

func deleteUserTestCases(t *testing.T) {
	for _, tc := range userTestCases.DeletUserTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestDeleteUser(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}
