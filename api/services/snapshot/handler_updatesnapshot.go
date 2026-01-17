package snapshot

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) UpdateSnapshot(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	snapshotPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateSnapshotPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot payload from request body"))
	}

	snapshotParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateSnapshotParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot params from request"))
	}

	existing, _ := h.snapshotStore.GetSnapshotByDate(c.Context(), snapshotPayload.Snap_date, userPayload.User_id)
	if existing.Snap_id != 0 && existing.Snap_id != snapshotParams.Id {
		return utils.SendError(c, fiber.StatusConflict, errors.New("snap_date must be unique"))
	}

	snapshot, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	if snapshot.Snap_id == 0 {
		return utils.SendError(c, fiber.StatusNotFound, errors.New("snapshot with the provided id does not exist"))
	}

	// Don't require value order to be passed in the payload; if it's missing, just use whatever the previous value was
	var valueOrder []int
	if snapshotPayload.Value_order != nil {
		valueOrder = snapshotPayload.Value_order
	} else {
		valueOrder = snapshot.Value_order
	}

	if len(valueOrder) > 0 {
		if code, err := h.validateSnapshotValueOrder(c, snapshot.Snap_id, snapshot.User_id, valueOrder); err != nil {
			return utils.SendError(c, code, err)
		}
	}

	snapshot, err = h.snapshotStore.UpdateSnapshot(c.Context(), &types.Snapshot{
		Snap_id:                 snapshot.Snap_id,
		Description:             snapshotPayload.Description,
		Snap_date:               snapshotPayload.Snap_date,
		Total:                   snapshot.Total,
		Weighted_er_pct:         snapshot.Weighted_er_pct,
		Rebalance_threshold_pct: snapshotPayload.Rebalance_threshold_pct,
		Value_order:             valueOrder,
		Benchmark_id:            snapshotPayload.Benchmark_id,
		User_id:                 snapshot.User_id,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"snapshot": snapshot})
}
