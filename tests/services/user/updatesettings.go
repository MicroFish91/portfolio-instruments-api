package user

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSettings(t *testing.T, route string, token string, userId int, payload any, expectedStatusCode int) {
	var updateSettingsResponse types.UpdateSettingsResponse
	var r string

	if route == "" {
		r = fmt.Sprintf("/api/v1/users/%d/settings", userId)
	} else {
		r = route
	}

	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, r, token, payload, &updateSettingsResponse)

	switch expectedStatusCode {
	case 200:
		p := payload.(user.UpdateSettingsPayload)
		assert.Equal(t, res.StatusCode, fiber.StatusOK)
		assert.EqualExportedValues(
			t,
			types.Settings{
				Settings_id:    updateSettingsResponse.Data.Settings.Settings_id,
				Reb_thresh_pct: p.Reb_thresh_pct,
				User_id:        userId,
				Created_at:     updateSettingsResponse.Data.Settings.Created_at,
				Updated_at:     updateSettingsResponse.Data.Settings.Updated_at,
			},
			updateSettingsResponse.Data.Settings,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
