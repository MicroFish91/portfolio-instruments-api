package holding

import (
	"errors"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *HoldingHandlerImpl) DeleteHolding(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	holdingParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteHoldingParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid holding params from request"))
	}

	holding, err := h.store.DeleteHolding(c.Context(), userPayload.User_id, holdingParams.Id)

	if err != nil {

		// If unable to delete because the holding is still being referenced in existing snapshots...
		isStillInUse := regexp.MustCompile(`(?i)violates\sforeign\skey\sconstraint`).Match([]byte(err.Error()))
		if isStillInUse {

			holding, err = h.store.GetHoldingById(c.Context(), userPayload.User_id, holdingParams.Id)
			if err != nil {
				return utils.SendError(c, utils.StatusCodeFromError(err), err)
			}

			if !holding.Is_deprecated {
				holding, err = h.store.UpdateHolding(c.Context(), &types.Holding{
					Holding_id:        holding.Holding_id,
					Name:              holding.Name,
					Ticker:            holding.Ticker,
					Asset_category:    holding.Asset_category,
					Expense_ratio_pct: holding.Expense_ratio_pct,
					Maturation_date:   holding.Maturation_date,
					Interest_rate_pct: holding.Interest_rate_pct,
					Is_deprecated:     true,
					User_id:           holding.User_id,
				})

				if err != nil {
					return utils.SendError(c, utils.StatusCodeFromError(err), err)
				}
			}

			return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
				"message": "Resource is still being referenced by an existing snapshot and could not be deleted. Resource marked as deprecated instead and will be deleted automatically when no longer in use.",
				"holding": holding,
			})
		}

		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message": "Resource deleted successfully.",
		"holding": holding,
	})
}
