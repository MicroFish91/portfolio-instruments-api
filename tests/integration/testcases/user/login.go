package user

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/gofiber/fiber/v3"
)

var LoginTestCases = []shared.PostTestCase{
	{
		Title: "201",
		Payload: auth.LoginPayload{
			Email:    "test_user@gmail.com",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusCreated,
	},
	{
		Title: "401",
		Payload: auth.LoginPayload{
			Email:    "test_user@gmail.com",
			Password: "fake-password",
		},
		ExpectedStatusCode: fiber.StatusUnauthorized,
	},
	{
		Title: "404",
		Payload: auth.LoginPayload{
			Email:    "test_user_fake@gmail.com",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusNotFound,
	},

	// -- Payload validation tests (400) --

	{
		Title: "400 No Email",
		Payload: auth.LoginPayload{
			Email:    "",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
	{
		Title: "400 Bad Email",
		Payload: auth.LoginPayload{
			Email:    "test_user",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
	{
		Title: "400 No Password",
		Payload: auth.LoginPayload{
			Email:    "test_user@gmail.com",
			Password: "",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
}
