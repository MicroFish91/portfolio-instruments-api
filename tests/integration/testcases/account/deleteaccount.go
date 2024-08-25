package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeleteAccountTests(t *testing.T, userId int, email string) []shared.DeleteTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.DeleteTestCase{
		{
			Title:              "200",
			ParameterId:        1,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        2,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        300,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
