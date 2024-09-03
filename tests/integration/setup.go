package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	authTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/auth"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func NewUserSetup(t *testing.T) (u types.User, tok string) {
	email := utils.GetRotatingEmail()
	password := "abcd1234"

	t.Run("Register", func(t2 *testing.T) {
		authTester.TestRegister(
			t2,
			auth.RegisterPayload{
				Email:    email,
				Password: password,
			},
			fiber.StatusCreated,
		)
	})

	var (
		testuser types.User
		token    string
	)

	t.Run("Login", func(t2 *testing.T) {
		testuser, token = authTester.TestLogin(
			t2,
			auth.LoginPayload{
				Email:    email,
				Password: password,
			},
			fiber.StatusCreated,
		)
	})

	return testuser, token
}
