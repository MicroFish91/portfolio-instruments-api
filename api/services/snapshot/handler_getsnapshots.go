package snapshot

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) GetSnapshots(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	queryPayload, ok := c.Locals(constants.LOCALS_REQ_QUERY).(GetSnapshotsQuery)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid query params from request"))
	}

	snapshots, pagination, err := h.snapshotStore.GetSnapshots(c.Context(), userPayload.User_id, types.GetSnapshotsStoreOptions{
		Snap_ids:        queryPayload.Snap_ids,
		Snap_date_lower: queryPayload.Snap_date_lower,
		Snap_date_upper: queryPayload.Snap_date_upper,
		Order_date_by:   queryPayload.Order_date_by,
		Current_page:    queryPayload.Current_page,
		Page_size:       queryPayload.Page_size,
	})

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"snapshots":  snapshots,
		"pagination": pagination,
	})
}
