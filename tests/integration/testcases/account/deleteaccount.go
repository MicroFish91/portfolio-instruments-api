package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeleteAccountTestCases(t *testing.T, accountId int, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:              "401",
			ParameterId:        accountId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "200",
			ParameterId:        accountId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "404",
			ParameterId:        accountId,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
