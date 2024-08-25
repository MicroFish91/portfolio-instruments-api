package user

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestGetSettings(t *testing.T, route string, token string, userId int, expectedStatusCode int) {
	var getSettingsResponse types.GetSettingsResponse
	var r string

	if route == "" {
		r = fmt.Sprintf("/api/v1/users/%d/settings", userId)
	} else {
		r = route
	}

	res := utils.SendGetRequest(t, r, token, &getSettingsResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		assert.EqualExportedValues(
			t,
			types.Settings{
				Settings_id:    getSettingsResponse.Data.Settings.Settings_id,
				Reb_thresh_pct: 10,
				User_id:        userId,
				Created_at:     getSettingsResponse.Data.Settings.Created_at,
				Updated_at:     getSettingsResponse.Data.Settings.Updated_at,
			},
			getSettingsResponse.Data.Settings,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
