package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetAccountTestCases(t *testing.T, accountId int, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		{
			Title:              "200",
			ParameterId:        accountId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        accountId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        9999,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
