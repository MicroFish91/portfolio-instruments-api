package usercases

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testcase"
	"github.com/gofiber/fiber/v3"
)

func GetSettingsTestCases(t *testing.T, userId int) []testcase.GetTestCase {
	tok401, err := auth.GenerateSignedJwt(userId, "test_user@gmail.com", "Default")
	if err != nil {
		t.Fatal(err)
	}
	tok401 = tok401[1:]

	tok403, err := auth.GenerateSignedJwt(100, "fake_user_100@gmail.com", "Default")
	if err != nil {
		t.Fatal(err)
	}

	return []testcase.GetTestCase{
		{
			Title:              "200",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        userId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "403 Token Id",
			ParameterId:        userId,
			ReplacementToken:   tok403,
			ExpectedStatusCode: fiber.StatusForbidden,
		},
		{
			Title:              "403 Param Id",
			Route:              "/api/v1/users/100/settings",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusForbidden,
		},

		// 400
		{
			Title:              "400 String Id",
			Route:              "/api/v1/users/test/settings",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Float Id",
			Route:              "/api/v1/users/1.0/settings",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Object Id",
			Route:              "/api/v1/users/{id:1}",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
