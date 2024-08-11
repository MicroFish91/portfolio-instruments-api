package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) DeleteUser(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user request body"))
	}

	userParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteUserParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user params from request"))
	}

	if userPayload.User_id != userParams.Id {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to delete user with the provided id"))
	}

	user, err := h.userStore.DeleteUser(c.Context(), userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message": "user deleted successfully",
		"user":    user,
	})
}
