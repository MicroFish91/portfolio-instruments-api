package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) UpdateSettings(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	settingsPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateSettingsPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid settings request body"))
	}

	if err := h.hasAuthorizedUserId(c, userPayload.User_id); err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, err)
	}

	if settingsPayload.Benchmark_id != 0 {
		benchmark, _ := h.benchmarkStore.GetBenchmarkById(c.Context(), userPayload.User_id, settingsPayload.Benchmark_id)
		if benchmark == nil {
			return utils.SendError(c, fiber.StatusConflict, errors.New("benchmark with provided id does not exist"))
		}
	}

	settings, err := h.userStore.UpdateSettings(c.Context(), &types.Settings{
		Reb_thresh_pct: settingsPayload.Reb_thresh_pct,
		Vp_thresh_pct:  settingsPayload.Vp_thresh_pct,
		Vp_enabled:     settingsPayload.Vp_enabled,
		Benchmark_id:   settingsPayload.Benchmark_id,
		User_id:        userPayload.User_id,
	})

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"settings": settings})
}
