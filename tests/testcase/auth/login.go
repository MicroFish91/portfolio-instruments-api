package auth

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	testUtils "github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T, p auth.LoginPayload, expectedStatusCode int) (u types.User, token string) {
	var loginResponse types.LoginResponse
	res := testUtils.SendAuthRequest(t, "/api/v1/login", &p, &loginResponse)

	switch expectedStatusCode {
	case 201:
		assert.Equal(t, fiber.StatusCreated, res.StatusCode)
		assert.NotEmpty(t, loginResponse.Data.Token)
		assert.EqualExportedValues(
			t,
			types.User{
				User_id:        loginResponse.Data.User.User_id,
				Email:          p.Email,
				User_role:      types.Default,
				Last_logged_in: loginResponse.Data.User.Last_logged_in,
				Created_at:     loginResponse.Data.User.Created_at,
				Updated_at:     loginResponse.Data.User.Updated_at,
			},
			loginResponse.Data.User,
		)
		return loginResponse.Data.User, loginResponse.Data.Token
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return types.User{}, ""
}
