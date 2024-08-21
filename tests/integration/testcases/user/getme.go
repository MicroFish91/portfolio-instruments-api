package user

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/gofiber/fiber/v3"
)

func GetMeTestCases(t *testing.T, userId int, email string) []shared.GetTestCase {
	tok401, err := auth.GenerateSignedJwt(userId, email, "Default")
	if err != nil {
		t.Fatal(err)
	}
	tok401 = tok401[1:]

	tok404, err := auth.GenerateSignedJwt(100, "fake_user_100@gmail.com", "Default")
	if err != nil {
		t.Fatal(err)
	}

	return []shared.GetTestCase{
		{
			Title:              "200",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ReplacementToken:   tok404,
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
