package account

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *AccountHandlerImpl) UpdateAccount(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	accountPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateAccountPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account request body"))
	}

	accountParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateAccountByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account params from request"))
	}

	account, err := h.store.UpdateAccount(c.Context(), types.Account{
		Account_id:    accountParams.Id,
		Name:          accountPayload.Name,
		Description:   accountPayload.Description,
		Tax_shelter:   accountPayload.Tax_shelter,
		Institution:   accountPayload.Institution,
		Is_deprecated: accountPayload.Is_deprecated,
		User_id:       userPayload.User_id,
	})

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"account": account})
}
