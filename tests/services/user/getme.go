package user

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetMe(t *testing.T, token string, userId int, expectedStatusCode int) {
	var getMeResponse types.GetMeResponse
	res := utils.SendGetRequest(t, "/api/v1/me", token, &getMeResponse)

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
					Email:          getMeResponse.Data.User.Email,
					User_role:      types.Default,
					Last_logged_in: getMeResponse.Data.User.Last_logged_in,
					Created_at:     getMeResponse.Data.User.Created_at,
					Updated_at:     getMeResponse.Data.User.Updated_at,
				},
				Settings: types.Settings{
					Settings_id:    getMeResponse.Data.Settings.Settings_id,
					Reb_thresh_pct: 10,
					User_id:        userId,
					Created_at:     getMeResponse.Data.User.Created_at,
					Updated_at:     getMeResponse.Data.User.Updated_at,
				},
			},
			getMeResponse.Data,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
