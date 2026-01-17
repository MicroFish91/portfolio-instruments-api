package snapshot

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) UpdateValueOrder(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	valueOrderPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateValueOrderPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid value_order payload from request body"))
	}

	snapshotParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateValueOrderParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot params from request"))
	}

	snapshot, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	if snapshot.Snap_id == 0 {
		return utils.SendError(c, fiber.StatusNotFound, errors.New("snapshot with the provided id does not exist"))
	}

	err = h.validateSnapshotValueOrder(c, snapshot.Snap_id, snapshot.User_id, valueOrderPayload.Value_order)
	if err != nil {
		return err
	}

	snapshot, err = h.snapshotStore.UpdateSnapshot(c.Context(), &types.Snapshot{
		Snap_id:                 snapshot.Snap_id,
		Description:             snapshot.Description,
		Snap_date:               snapshot.Snap_date,
		Total:                   snapshot.Total,
		Weighted_er_pct:         snapshot.Weighted_er_pct,
		Rebalance_threshold_pct: snapshot.Rebalance_threshold_pct,
		Value_order:             valueOrderPayload.Value_order,
		Benchmark_id:            snapshot.Benchmark_id,
		User_id:                 snapshot.User_id,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"snapshot": snapshot})
}

func (h *SnapshotHandlerImpl) validateSnapshotValueOrder(c fiber.Ctx, snapshotId, userId int, valueOrder []int) error {
	if len(valueOrder) == 0 {
		return utils.SendError(c, fiber.StatusInternalServerError, errors.New("internal error: value_order is required for validation"))
	}

	snapshotValues, err := h.snapshotValueStore.GetSnapshotValues(c.Context(), snapshotId, userId)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// Build a set of valid snapshot value IDs
	validIds := make(map[int]bool, len(snapshotValues))
	for _, sv := range snapshotValues {
		validIds[sv.Snap_val_id] = true
	}

	// Check for duplicates and invalid IDs
	seen := make(map[int]bool, len(valueOrder))
	for _, id := range valueOrder {
		if seen[id] {
			return utils.SendError(c, fiber.StatusBadRequest, errors.New("value_order contains duplicates"))
		}
		if !validIds[id] {
			return utils.SendError(c, fiber.StatusBadRequest, errors.New("value_order contains invalid snapshot value IDs"))
		}
		seen[id] = true
	}

	// Ensure all snapshot values are accounted for
	if len(valueOrder) != len(snapshotValues) {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("value_order must include all snapshot value IDs"))
	}

	return nil
}
