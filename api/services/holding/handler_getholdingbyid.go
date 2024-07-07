package holding

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *HoldingHandlerImpl) GetHoldingById(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	holdingParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(GetHoldingByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("uanble to parse valid holding params from request"))
	}

	holding, err := h.store.GetHoldingById(userPayload.User_id, holdingParams.Id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"holding": holding})
}
