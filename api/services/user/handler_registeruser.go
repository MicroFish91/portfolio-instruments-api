package user

import (
	"context"
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *UserHandlerImpl) RegisterUser(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(RegisterUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	getUserCtx, getUserCancel := context.WithTimeout(c.Context(), time.Second*5)
	defer getUserCancel()

	user, _ := h.store.GetUserByEmail(getUserCtx, userPayload.Email)
	if user != nil {
		return utils.SendError(c, fiber.StatusConflict, errors.New("user with provided email already exists"))
	}

	encpw, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	registerUserCtx, registerUserCancel := context.WithTimeout(c.Context(), time.Second*5)
	defer registerUserCancel()

	err = h.store.RegisterUser(
		registerUserCtx,
		&types.User{
			Email:        userPayload.Email,
			Enc_password: encpw,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"message": "user registered successfully"})
}
