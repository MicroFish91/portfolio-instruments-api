package holding

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *HoldingHandlerImpl) GetHoldings(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	queryPayload, ok := c.Locals(constants.LOCALS_REQ_QUERY).(GetHoldingsQuery)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid query params from request"))
	}

	holdings, pagination, err := h.store.GetHoldings(userPayload.User_id, &types.GetHoldingsStoreOptions{
		Holding_ids:    queryPayload.Holding_ids,
		Ticker:         queryPayload.Ticker,
		Asset_category: queryPayload.Asset_category,
		Is_deprecated:  queryPayload.Is_deprecated,
		Current_page:   queryPayload.Current_page,
		Page_size:      queryPayload.Page_size,
	})

	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"holdings":   holdings,
		"pagination": pagination,
	})
}
