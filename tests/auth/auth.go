package auth

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
)

func TestAuth(t *testing.T) {
	t.Run("POST://api/v1/register", func(t *testing.T) {
		TestRegister(
			t,
			auth.RegisterPayload{
				Email:    "test_user@gmail.com",
				Password: "abcd1234",
			},
			TestRegisterOptions{Parallel: false},
		)
	})
	t.Run("POST://api/v1/login", func(t *testing.T) {
		TestLogin(
			t,
			auth.LoginPayload{
				Email:    "test_user@gmail.com",
				Password: "abcd1234",
			},
			LoginTestOptions{Parallel: false},
		)
	})
}
