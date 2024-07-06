package user

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) LoginUser(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(LoginUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	user, err := h.store.GetUserByEmail(userPayload.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	if err := auth.CompareHashAndPassword(user.Enc_password, userPayload.Password); err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("invalid login credentials"))
	}

	jwt, err := auth.GenerateSignedJwt(user.User_id, user.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"token": jwt,
		"user":  user,
	})
}
