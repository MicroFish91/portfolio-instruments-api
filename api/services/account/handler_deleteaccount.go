package account

import (
	"errors"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *AccountHandlerImpl) DeleteAccount(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	accountParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteAccountByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account params from request"))
	}

	account, err := h.store.DeleteAccount(c.Context(), userPayload.User_id, accountParams.Id)

	if err != nil {

		// If unable to delete because the holding is still being referenced in existing snapshots...
		isStillInUse := regexp.MustCompile(`(?i)violates\sforeign\skey\sconstraint`).Match([]byte(err.Error()))
		if isStillInUse {

			account, err = h.store.GetAccountById(c.Context(), userPayload.User_id, accountParams.Id)
			if err != nil {
				return utils.SendError(c, utils.StatusCodeFromError(err), err)
			}

			if !account.Is_deprecated {
				account, err = h.store.UpdateAccount(c.Context(), &types.Account{
					Account_id:    accountParams.Id,
					Name:          account.Name,
					Description:   account.Description,
					Tax_shelter:   account.Tax_shelter,
					Institution:   account.Institution,
					Is_deprecated: true,
					User_id:       userPayload.User_id,
				})

				if err != nil {
					return utils.SendError(c, utils.StatusCodeFromError(err), err)
				}
			}

			return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
				"message": "Resource is still being referenced by an existing snapshot and could not be deleted. Resource marked as deprecated instead and will be deleted automatically when no longer in use.",
				"account": account,
			})
		}

		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message": "Resource deleted successfully.",
		"account": account,
	})
}
