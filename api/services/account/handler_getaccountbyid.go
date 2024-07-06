package account

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *AccountHandlerImpl) GetAccountById(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	accountParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(GetAccountByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account params from request"))
	}

	account, err := h.store.GetAccountById(userPayload.User_id, accountParams.Id)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"account": account})
}
