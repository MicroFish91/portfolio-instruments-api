package holding

import (
	"context"
	"errors"
	"time"

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
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid holding params from request"))
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
	defer cancel()

	holding, err := h.store.GetHoldingById(ctx, userPayload.User_id, holdingParams.Id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"holding": holding})
}
