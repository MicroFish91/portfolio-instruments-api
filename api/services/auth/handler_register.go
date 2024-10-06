package auth

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *AuthHandlerImpl) Register(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(RegisterPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user request body"))
	}

	u, _ := h.store.GetUserByEmail(c.Context(), userPayload.Email)
	if u.User_id != 0 {
		return utils.SendError(c, fiber.StatusConflict, errors.New("user with provided email already exists"))
	}

	encpw, err := hashPassword(userPayload.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	user, err := h.store.CreateUser(
		c.Context(),
		&types.User{
			Email:        userPayload.Email,
			Enc_password: encpw,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	settings, err := h.store.CreateSettings(c.Context(), &types.Settings{
		Reb_thresh_pct: 10,
		User_id:        user.User_id,
	})

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{
		"user":     user,
		"settings": settings,
	})
}
