package user

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetUserById(t *testing.T, route string, token string, userId int, expectedStatusCode int) {
	var getUserResponse types.GetUserByIdResponse
	var r string

	if route == "" {
		r = fmt.Sprintf("/api/v2/users/%d", userId)
	} else {
		r = route
	}

	res := utils.SendGetRequest(t, r, token, &getUserResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		assert.EqualExportedValues(
			t,
			struct {
				User types.User `json:"user"`
			}{
				User: types.User{
					User_id:        userId,
					Email:          getUserResponse.Data.User.Email,
					User_role:      types.Default,
					Last_logged_in: getUserResponse.Data.User.Last_logged_in,
					Created_at:     getUserResponse.Data.User.Created_at,
					Updated_at:     getUserResponse.Data.User.Updated_at,
				},
			},
			getUserResponse.Data,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
