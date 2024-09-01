package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeleteBenchmarkTestCases(t *testing.T, benchmarkId int, userId int, email string) []shared.DeleteTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.DeleteTestCase{
		{
			Title:              "401",
			ParameterId:        benchmarkId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "200",
			ParameterId:        benchmarkId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "404",
			ParameterId:        benchmarkId,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
