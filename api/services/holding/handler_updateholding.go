package holding

import (
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *HoldingHandlerImpl) UpdateHolding(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	holdingPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateHoldingPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid holding payload from request body"))
	}

	holdingParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateHoldingParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid holding params from request"))
	}

	// Ensure only one unique "active" ticker per user
	if holdingPayload.Ticker != "" && !holdingPayload.Is_deprecated {
		existingHolding, _ := h.store.GetHoldingByTicker(c.Context(), holdingPayload.Ticker, userPayload.User_id)
		if existingHolding.Holding_id != 0 {
			return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`user already has holding with ticker symbol "%s"`, existingHolding.Ticker))
		}
	}

	holding, err := h.store.UpdateHolding(
		c.Context(),
		types.Holding{
			Holding_id:        holdingParams.Id,
			Name:              holdingPayload.Name,
			Ticker:            holdingPayload.Ticker,
			Asset_category:    holdingPayload.Asset_category,
			Expense_ratio_pct: holdingPayload.Expense_ratio_pct,
			Maturation_date:   holdingPayload.Maturation_date,
			Interest_rate_pct: holdingPayload.Interest_rate_pct,
			Is_deprecated:     holdingPayload.Is_deprecated,
			User_id:           userPayload.User_id,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"holding": holding})
}
