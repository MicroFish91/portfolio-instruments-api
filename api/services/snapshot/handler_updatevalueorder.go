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
