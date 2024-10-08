package account

import (
	"errors"
	"fmt"

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

	// Ensure account name is unique per user
	existingAccount, _ := h.store.GetAccountByName(c.Context(), accountPayload.Name, userPayload.User_id)
	if existingAccount.Account_id != 0 {
		return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`user already has account with name "%s"`, existingAccount.Name))
	}

	account, err := h.store.CreateAccount(
		c.Context(),
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
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"account": account})
}
