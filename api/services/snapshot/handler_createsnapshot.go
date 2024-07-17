package snapshot

import (
	"context"
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) CreateSnapshot(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	snapshotPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateSnapshotPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot payload from request body"))
	}

	// Todo: Should we check that all accounts and holdings exist first?

	createSnapshotCtx, cancelSnapshot := context.WithTimeout(c.Context(), time.Second*5)
	defer cancelSnapshot()

	snapshot, err := h.store.CreateSnapshot(
		createSnapshotCtx,
		&types.Snapshot{
			Snap_date: snapshotPayload.Snap_date,
			User_id:   userPayload.User_id,
		},
	)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	createSnapshotValsCtx, cancelSnapshotVals := context.WithTimeout(createSnapshotCtx, time.Second*5)
	defer cancelSnapshotVals()

	var snapshotValues []types.SnapshotValues
	for _, sv := range snapshotPayload.Snapshot_values {
		snapshotVal, err := h.store.CreateSnapshotValues(
			createSnapshotValsCtx,
			&types.SnapshotValues{
				Snap_id:        snapshot.Snap_id,
				Account_id:     sv.Account_id,
				Holding_id:     sv.Holding_id,
				Total:          sv.Total,
				Skip_rebalance: sv.Skip_rebalance,
				User_id:        userPayload.User_id,
			},
		)
		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}
		snapshotValues = append(snapshotValues, *snapshotVal)
	}

	refreshSnapshotCtx, cancelRefresh := context.WithTimeout(createSnapshotValsCtx, time.Second*10)
	defer cancelRefresh()

	total, err := h.store.RefreshSnapshotTotal(refreshSnapshotCtx, userPayload.User_id, snapshot.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	snapshot.Total = total
	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{
		"snapshot":        snapshot,
		"snapshot_values": snapshotValues,
	})
}
