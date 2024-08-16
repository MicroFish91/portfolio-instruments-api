package usercases

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testcases"
	"github.com/gofiber/fiber/v3"
)

var RegisterTestCases = []testcases.TestCase{
	// 201
	{
		Title: "201",
		Payload: auth.RegisterPayload{
			Email:    "test_user@gmail.com",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusCreated,
	},

	// 409
	{
		Title: "409",
		Payload: auth.RegisterPayload{
			Email:    "test_user@gmail.com",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusConflict,
	},

	// -- Payload validation tests (400) --

	{
		Title: "400 No Email",
		Payload: auth.RegisterPayload{
			Email:    "",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
	{
		Title: "400 Bad Email",
		Payload: auth.RegisterPayload{
			Email:    "test_user",
			Password: "abcd1234",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
	{
		Title: "400 No Password",
		Payload: auth.RegisterPayload{
			Email:    "test_user@gmail.com",
			Password: "",
		},
		ExpectedStatusCode: fiber.StatusBadRequest,
	},
}
