package user

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeletUserTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, tok403, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
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
			Route:              "/api/v1/users/100",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusForbidden,
		},

		// 400
		{
			Title:              "400 String Id",
			Route:              "/api/v1/users/test",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Float Id",
			Route:              "/api/v1/users/1.0",
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
