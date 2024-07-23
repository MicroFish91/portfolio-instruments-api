package snapshot

import (
	"errors"
	"fmt"

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

	// Todo: If benchmark_id provided, ensure that it already exists

	// Verify accounts exist
	accIdsSet := map[int]struct{}{}
	for _, svpayload := range snapshotPayload.Snapshot_values {
		if _, checked := accIdsSet[svpayload.Account_id]; checked {
			continue
		}

		acc, _ := h.accountStore.GetAccountById(c.Context(), userPayload.User_id, svpayload.Account_id)
		if acc == nil {
			return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`specified account with id "%d" does not exist`, svpayload.Account_id))
		}
		accIdsSet[svpayload.Account_id] = struct{}{}
	}

	// Verify holdings exist
	holdIdsSet := map[int]struct{}{}
	for _, svpayload := range snapshotPayload.Snapshot_values {
		if _, checked := holdIdsSet[svpayload.Holding_id]; checked {
			continue
		}

		hold, _ := h.holdingStore.GetHoldingById(c.Context(), userPayload.User_id, svpayload.Holding_id)
		if hold == nil {
			return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`specified holding with id "%d" does not exist`, svpayload.Holding_id))
		}
		holdIdsSet[svpayload.Holding_id] = struct{}{}
	}

	// Create snapshot
	snapshot, err := h.snapshotStore.CreateSnapshot(
		c.Context(),
		&types.Snapshot{
			Snap_date:    snapshotPayload.Snap_date,
			Description:  snapshotPayload.Description,
			User_id:      userPayload.User_id,
			Benchmark_id: snapshotPayload.Benchmark_id,
		},
	)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// Create snapshot values
	var snapshotValues []types.SnapshotValues
	for _, sv := range snapshotPayload.Snapshot_values {
		snapshotVal, err := h.snapshotStore.CreateSnapshotValues(
			c.Context(),
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

	// Acquire snapshot total
	total, err := h.snapshotStore.RefreshSnapshotTotal(c.Context(), userPayload.User_id, snapshot.Snap_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	snapshot.Total = total

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{
		"snapshot":        snapshot,
		"snapshot_values": snapshotValues,
	})
}
