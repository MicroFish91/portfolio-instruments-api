package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeleteHoldingTestCases(t *testing.T, holdingId int, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:              "401",
			ParameterId:        holdingId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "200",
			ParameterId:        holdingId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "404",
			ParameterId:        holdingId,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
