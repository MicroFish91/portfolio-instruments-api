package account

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

func (h *AccountHandlerImpl) CreateAccount(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	accountPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateAccountPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account request body"))
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
	defer cancel()

	err := h.store.CreateAccount(
		ctx,
		&types.Account{
			Name:          accountPayload.Name,
			Description:   accountPayload.Description,
			Tax_shelter:   accountPayload.Tax_shelter,
			Institution:   accountPayload.Institution,
			Is_deprecated: accountPayload.Is_deprecated,
			User_id:       userPayload.User_id,
		},
	)

	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}
	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{})
}
