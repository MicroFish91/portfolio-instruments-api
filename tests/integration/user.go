package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	userTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/user"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/auth"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
)

var (
	us_token    string
	us_testuser types.User

	email    string = utils.GetRotatingEmail()
	password string = "abcd1234"
)

func TestUserService(t *testing.T) {
	t.Run("POST://api/v2/register", registerTests)
	t.Run("POST://api/v2/login", loginTests)
	t.Run("GET://api/v2/me", getMeTests)
	t.Run("GET://api/v2/users/:id", getUserByIdTests)
	t.Run("DEL://api/v2/users/:id", deleteUserTests)
}

func registerTests(t *testing.T) {
	for _, tc := range userTestCases.RegisterTestCases(email, password) {
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

func loginTests(t *testing.T) {
	for _, tc := range userTestCases.LoginTestCases(email, password) {
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

func getMeTests(t *testing.T) {
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

func getUserByIdTests(t *testing.T) {
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

func deleteUserTests(t *testing.T) {
	for _, tc := range userTestCases.DeleteUserTestCases(t, us_testuser.User_id, us_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := us_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}
			userTester.TestDeleteUser(t2, tc.Route, tok, tc.ParameterId, tc.ExpectedStatusCode)
		})
	}
}
