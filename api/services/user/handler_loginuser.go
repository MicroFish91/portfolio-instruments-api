package user

import (
	"context"
	"errors"
	"time"

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

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
	defer cancel()

	user, err := h.store.GetUserByEmail(ctx, userPayload.Email)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
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
