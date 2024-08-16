package auth

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	testUtils "github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T, p auth.RegisterPayload, expectedStatusCode int) {
	var responseBody types.RegisterResponse
	res := testUtils.SendPostRequest(t, "/api/v1/register", &p, &responseBody)
	defer res.Body.Close()

	switch expectedStatusCode {
	case 201:
		assert.Equal(t, res.StatusCode, fiber.StatusCreated)
		assert.Equal(t, p.Email, responseBody.Data.User.Email)
	case 400:
		assert.Equal(t, res.StatusCode, fiber.StatusBadRequest)
	case 409:
		assert.Equal(t, res.StatusCode, fiber.StatusConflict)
	default:
		t.Fatal("provided anunexpected status code")
	}
}
