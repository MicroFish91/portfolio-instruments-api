package snapshotvalue

import (
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotValueHandlerImpl) DeleteSnapshotValue(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	svParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteSnapshotValueParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot_value params from request"))
	}

	// snapshot_value
	snapshotvalue, err := h.snapshotValueStore.DeleteSnapshotValue(c.Context(), svParams.Snap_id, svParams.Snap_val_id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// Refresh snapshot total
	total, err := h.snapshotStore.RefreshSnapshotTotal(c.Context(), userPayload.User_id, svParams.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed to refresh snapshot total: %s", err.Error()))
	}

	// Refresh snapshot weighted_er
	er, err := h.snapshotStore.RefreshSnapshotWeightedER(c.Context(), snapshotvalue.User_id, svParams.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed to refresh snapshot weighted_er: %s", err.Error()))
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message":             "resource deleted successfully",
		"snapshot_value":      snapshotvalue,
		"snapshot_total":      total,
		"snapshot_weighteder": er,
	})
}
