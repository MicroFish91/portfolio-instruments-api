package auth

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	testUtils "github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T, p auth.LoginPayload, expectedStatusCode int) {
	var loginResponse types.LoginResponse
	res := testUtils.SendPostRequest(t, "/api/v1/login", &p, &loginResponse)
	defer res.Body.Close()

	switch expectedStatusCode {
	case 201:
		assert.Equal(t, res.StatusCode, fiber.StatusCreated)
		assert.Equal(t, p.Email, loginResponse.Data.User.Email)
	case 400:
		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
	case 401:
		assert.Equal(t, res.StatusCode, fiber.StatusUnauthorized)
	case 404:
		assert.Equal(t, res.StatusCode, fiber.StatusNotFound)
	case 409:
		assert.Equal(t, res.StatusCode, fiber.StatusConflict)
	default:
		t.Fatal("provided an unexpected status code")
	}
}
