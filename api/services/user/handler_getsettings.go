package user

import (
	"errors"
	"strconv"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) GetSettings(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	uid := c.Params("id", "")
	if uid != "" {
		userId, err := strconv.Atoi(uid)
		if err != nil {
			return utils.SendError(c, fiber.StatusBadRequest, errors.New("user id must be an integer"))
		}

		if userId != userPayload.User_id {
			return utils.SendError(c, fiber.StatusUnauthorized, errors.New("provided token does not correspond with the requested user"))
		}
	}

	settings, err := h.store.GetSettings(c.Context(), userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"settings": settings})
}
