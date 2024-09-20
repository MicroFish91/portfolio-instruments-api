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
	if err := h.verifySnapshotById(c, svParams.Snap_id, userPayload.User_id); err != nil {
		return utils.SendError(c, fiber.StatusNotFound, err)
	}

	// snapshotvalue
	snapshotvalue, err := h.snapshotValueStore.CreateSnapshotValue(c.Context(), types.SnapshotValue{
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

	// Refresh snapshot total
	if _, err = h.snapshotStore.RefreshSnapshotTotal(c.Context(), userPayload.User_id, svParams.Snap_id); err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed to refresh snapshot total: %s", err.Error()))
	}

	// Refresh snapshot weighted_er
	if _, err = h.snapshotStore.RefreshSnapshotWeightedER(c.Context(), snapshotvalue.User_id, svParams.Snap_id); err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), fmt.Errorf("failed to refresh snapshot weighted_er: %s", err.Error()))
	}

	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"snapshot_value": snapshotvalue})
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

func (h *SnapshotValueHandlerImpl) verifySnapshotById(c fiber.Ctx, snapshotId, userId int) error {
	_, _, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotId, userId)
	if err != nil {
		return fmt.Errorf(`specified snapshot with id "%d" not found for the current user`, snapshotId)
	}
	return nil
}
