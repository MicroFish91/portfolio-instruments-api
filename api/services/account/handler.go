package account

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

type AccountHandlerImpl struct {
	store types.AccountStore
}

func NewAccountHandler(store types.AccountStore) *AccountHandlerImpl {
	return &AccountHandlerImpl{
		store: store,
	}
}

func (h *AccountHandlerImpl) HandleCreateAccount(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	accountPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateAccountPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account request body"))
	}

	err := h.store.CreateAccount(&types.Account{
		Name:        accountPayload.Name,
		Description: accountPayload.Description,
		Tax_Shelter: accountPayload.Shelter_type,
		Institution: accountPayload.Institution,
		User_id:     userPayload.User_id,
	})

	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{})
}
