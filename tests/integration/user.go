package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	userTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/user"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/auth"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/services/user"
)

var token string
var testUser types.User

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
	for _, tc := range userTestCases.RegisterTestCases {
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
	for _, tc := range userTestCases.LoginTestCases {
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
			if token == "" && tok != "" {
				token = tok
				testUser = u
			}
		})
	}
}

func getMeTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetMeTestCases(t, testUser.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetMe(t2, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func getUserByIdTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetUserByIdTestCases(t, testUser.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetUserById(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func getSettingsTestCases(t *testing.T) {
	for _, tc := range userTestCases.GetSettingsTestCases(t, testUser.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestGetSettings(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}

func updateSettingsTestCases(t *testing.T) {
	for _, tc := range userTestCases.UpdateSettingsTestCases(t, testUser.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestUpdateSettings(t2, tc.Route, tok, tc.ParameterId, tc.Payload, tc.ExpectedStatusCode)
		})
	}
}

func deleteUserTestCases(t *testing.T) {
	for _, tc := range userTestCases.DeletUserTestCases(t, testUser.User_id) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestDeleteUser(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}
