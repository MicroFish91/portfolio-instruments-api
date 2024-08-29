package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetHoldingTestCases(t *testing.T, holdId, userId int, email string) []shared.GetTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.GetTestCase{
		{
			Title:              "200",
			ParameterId:        holdId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        holdId,
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
