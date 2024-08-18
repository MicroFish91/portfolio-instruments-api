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
		r = fmt.Sprintf("/api/v1/users/%d", userId)
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
				User     types.User     `json:"user"`
				Settings types.Settings `json:"settings"`
			}{
				User: types.User{
					User_id:        userId,
					Email:          getUserResponse.Data.User.Email,
					User_role:      types.Default,
					Last_logged_in: getUserResponse.Data.User.Last_logged_in,
					Created_at:     getUserResponse.Data.User.Created_at,
					Updated_at:     getUserResponse.Data.User.Updated_at,
				},
				Settings: types.Settings{
					Settings_id:    getUserResponse.Data.Settings.Settings_id,
					Reb_thresh_pct: 10,
					User_id:        userId,
					Created_at:     getUserResponse.Data.User.Created_at,
					Updated_at:     getUserResponse.Data.User.Updated_at,
				},
			},
			getUserResponse.Data,
		)
	case 400:
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	case 401:
		assert.Equal(t, fiber.StatusUnauthorized, res.StatusCode)
	case 403:
		assert.Equal(t, fiber.StatusForbidden, res.StatusCode)
	default:
		t.Fatal("provided an unexpected status code")
	}
}
