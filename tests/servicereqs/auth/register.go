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
	var registerResponse types.RegisterResponse
	res := testUtils.SendAuthRequest(t, "/api/v2/register", &p, &registerResponse)

	switch expectedStatusCode {
	case 201:
		assert.Equal(t, fiber.StatusCreated, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.User{
				User_id:        registerResponse.Data.User.User_id,
				Email:          p.Email,
				User_role:      types.Default,
				Last_logged_in: registerResponse.Data.User.Last_logged_in,
				Created_at:     registerResponse.Data.User.Created_at,
				Updated_at:     registerResponse.Data.User.Updated_at,
			},
			registerResponse.Data.User,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
