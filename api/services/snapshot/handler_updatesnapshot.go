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

	snapshot, _, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	if snapshot.Snap_id == 0 {
		return utils.SendError(c, fiber.StatusNotFound, errors.New("snapshot with the provided id does not exist"))
	}

	snapshot, err = h.snapshotStore.UpdateSnapshot(c.Context(), types.Snapshot{
		Snap_id:      snapshot.Snap_id,
		Description:  snapshotPayload.Description,
		Snap_date:    snapshotPayload.Snap_date,
		Total:        snapshot.Total,
		Weighted_er:  snapshot.Weighted_er,
		Benchmark_id: snapshotPayload.Benchmark_id,
		User_id:      snapshot.User_id,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"snapshot": snapshot})
}
