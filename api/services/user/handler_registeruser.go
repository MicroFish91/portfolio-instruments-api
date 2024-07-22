package user

import (
	"errors"

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

	u, _ := h.userStore.GetUserByEmail(c.Context(), userPayload.Email)
	if u != nil {
		return utils.SendError(c, fiber.StatusConflict, errors.New("user with provided email already exists"))
	}

	encpw, err := auth.HashPassword(userPayload.Password)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	user, err := h.userStore.RegisterUser(
		c.Context(),
		&types.User{
			Email:        userPayload.Email,
			Enc_password: encpw,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	settings, err := h.settingsStore.CreateSettings(c.Context(), &types.Settings{
		Reb_thresh_pct: 10,
		Vp_thresh_pct:  0,
		Vp_enabled:     false,
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
