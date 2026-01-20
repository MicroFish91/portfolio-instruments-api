package snapshotvalue

import (
	"errors"
	"fmt"
	"slices"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
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
	snapshotValue, err := h.snapshotValueStore.DeleteSnapshotValue(c.Context(), svParams.Snap_id, svParams.Snap_val_id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// If a value_order exists, append the newly created snapshot_value to the end of it
	snapshot, err := h.snapshotStore.GetSnapshotById(c.Context(), svParams.Snap_id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to update snapshot value_order: %s", err.Error()))
	}

	if len(snapshot.Value_order) > 0 {
		valueOrder := slices.DeleteFunc(snapshot.Value_order, func(valId int) bool {
			return valId == snapshotValue.Snap_val_id
		})

		_, err := h.snapshotStore.UpdateSnapshot(c.Context(), &types.Snapshot{
			Snap_id:                 snapshot.Snap_id,
			Description:             snapshot.Description,
			Snap_date:               snapshot.Snap_date,
			Total:                   snapshot.Total,
			Weighted_er_pct:         snapshot.Weighted_er_pct,
			Rebalance_threshold_pct: snapshot.Rebalance_threshold_pct,
			Value_order:             valueOrder,
			Benchmark_id:            snapshot.Benchmark_id,
			User_id:                 snapshot.User_id,
		})

		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to update snapshot value_order: %s", err.Error()))
		}
	}

	// Refresh snapshot total
	total, err := h.snapshotStore.RefreshSnapshotTotal(c.Context(), userPayload.User_id, svParams.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to refresh snapshot total: %s", err.Error()))
	}

	// Refresh snapshot weighted_er
	er, err := h.snapshotStore.RefreshSnapshotWeightedER(c.Context(), snapshotValue.User_id, svParams.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to refresh snapshot weighted_er: %s", err.Error()))
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message":             "resource deleted successfully",
		"snapshot_value":      snapshotValue,
		"snapshot_total":      total,
		"snapshot_weighteder": er,
	})
}
