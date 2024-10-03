package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) UpdateVerification(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	userParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateVerificationParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user params from request"))
	}

	if userPayload.User_role != types.Admin {
		return utils.SendError(c, fiber.StatusForbidden, errors.New("you must be an admin to access this route"))
	}

	user, err := h.userStore.UpdateVerification(c.Context(), userParams.Id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"user": user})
}
