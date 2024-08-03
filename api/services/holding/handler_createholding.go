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

func (h *HoldingHandlerImpl) CreateHolding(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	holdingPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateHoldingPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid holding payload from request body"))
	}

	// Ensure ticker holdings are unique per user
	if holdingPayload.Ticker != "" {
		existingHolding, _ := h.store.GetHoldingByTicker(c.Context(), holdingPayload.Ticker, userPayload.User_id)
		if existingHolding.Holding_id != 0 {
			return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`user already has holding with ticker symbol "%s"`, existingHolding.Ticker))
		}
	}

	holding, err := h.store.CreateHolding(
		c.Context(),
		types.Holding{
			Name:            holdingPayload.Name,
			Ticker:          holdingPayload.Ticker,
			Asset_category:  holdingPayload.Asset_category,
			Expense_ratio:   holdingPayload.Expense_ratio,
			Maturation_date: holdingPayload.Maturation_date,
			Interest_rate:   holdingPayload.Interest_rate,
			Is_deprecated:   holdingPayload.Is_deprecated,
			User_id:         userPayload.User_id,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"holding": holding})
}
