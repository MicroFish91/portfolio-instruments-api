package snapshotvalue

import (
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotValueHandlerImpl) CreateSnapshotValue(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	svPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateSnapshotValuePayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot_value payload from request body"))
	}

	svParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(CreateSnapshotValueParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot_value params from request"))
	}

	// Verify account
	if err := h.verifyAccountById(c, svPayload.Account_id, userPayload.User_id); err != nil {
		return utils.SendError(c, fiber.StatusNotFound, err)
	}

	// Verify holding
	if err := h.verifyHoldingById(c, svPayload.Holding_id, userPayload.User_id); err != nil {
		return utils.SendError(c, fiber.StatusNotFound, err)
	}

	// Verify snapshot
	snapshot, err := h.verifySnapshotById(c, svParams.Snap_id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, err)
	}

	// snapshotValue
	snapshotValue, err := h.snapshotValueStore.CreateSnapshotValue(c.Context(), &types.SnapshotValue{
		Snap_id:        svParams.Snap_id,
		Account_id:     svPayload.Account_id,
		Holding_id:     svPayload.Holding_id,
		Total:          svPayload.Total,
		Skip_rebalance: svPayload.Skip_rebalance,
		User_id:        userPayload.User_id,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// If a value_order exists, append the newly created snapshot_value to the end of it
	if len(snapshot.Value_order) > 0 {
		valueOrder := append(snapshot.Value_order, snapshotValue.Snap_val_id)
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
	if _, err = h.snapshotStore.RefreshSnapshotTotal(c.Context(), userPayload.User_id, svParams.Snap_id); err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to refresh snapshot total: %s", err.Error()))
	}

	// Refresh snapshot weighted_er
	if _, err = h.snapshotStore.RefreshSnapshotWeightedER(c.Context(), snapshotValue.User_id, svParams.Snap_id); err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed while attempting to refresh snapshot weighted_er: %s", err.Error()))
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"snapshot_value": snapshotValue})
}

func (h *SnapshotValueHandlerImpl) verifyAccountById(c fiber.Ctx, accountId, userId int) error {
	_, err := h.accountStore.GetAccountById(c.Context(), userId, accountId)
	if err != nil {
		return fmt.Errorf(`specified account with id "%d" not found for the current user`, accountId)
	}
	return nil
}

func (h *SnapshotValueHandlerImpl) verifyHoldingById(c fiber.Ctx, holdingId, userId int) error {
	_, err := h.holdingStore.GetHoldingById(c.Context(), userId, holdingId)
	if err != nil {
		return fmt.Errorf(`specified holding with id "%d" not found for the current user`, holdingId)
	}
	return nil
}

func (h *SnapshotValueHandlerImpl) verifySnapshotById(c fiber.Ctx, snapshotId, userId int) (types.Snapshot, error) {
	snapshot, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotId, userId)
	if err != nil {
		return types.Snapshot{}, fmt.Errorf(`specified snapshot with id "%d" not found for the current user`, snapshotId)
	}
	return snapshot, nil
}
