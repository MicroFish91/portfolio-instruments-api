package user

import (
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/gofiber/fiber/v3"
)

func GetRegisterTestCases(email string, password string) []shared.TestCase {
	return []shared.TestCase{
		{
			Title: "201",
			Payload: auth.RegisterPayload{
				Email:    email,
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "409",
			Payload: auth.RegisterPayload{
				Email:    strings.ToUpper(email[:1]) + email[1:],
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// -- Payload validation tests (400) --

		{
			Title: "400 No Email",
			Payload: auth.RegisterPayload{
				Email:    "",
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Email",
			Payload: auth.RegisterPayload{
				Email:    "test_user",
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Password",
			Payload: auth.RegisterPayload{
				Email:    email,
				Password: "",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
