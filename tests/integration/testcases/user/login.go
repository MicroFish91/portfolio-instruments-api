package user

import (
	"strings"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/gofiber/fiber/v3"
)

func LoginTestCases(email string, password string) []testcases.TestCase {
	return []testcases.TestCase{
		{
			Title: "201",
			Payload: auth.LoginPayload{
				Email:    email,
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Case Insensitive",
			Payload: auth.LoginPayload{
				Email:    strings.ToUpper(email[:1]) + email[1:],
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "401",
			Payload: auth.LoginPayload{
				Email:    email,
				Password: "fake-pass",
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title: "404",
			Payload: auth.LoginPayload{
				Email:    "test_user_fake@gmail.com",
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// -- Payload validation tests (400) --

		{
			Title: "400 No Email",
			Payload: auth.LoginPayload{
				Email:    "",
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Email",
			Payload: auth.LoginPayload{
				Email:    "test_user",
				Password: password,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Password",
			Payload: auth.LoginPayload{
				Email:    email,
				Password: "",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
