package snapshotvalue

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotValueHandlerImpl) UpdateSnapshotValue(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid user from token"))
	}

	svPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateSnapshotValuePayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot_value payload from request body"))
	}

	svParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateSnapshotValueParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot_value params from request"))
	}

	snapshotValue, err := h.snapshotValueStore.UpdateSnapshotValue(c.Context(), types.SnapshotValue{
		Snap_val_id:    svParams.Snap_val_id,
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

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"snapshot_value": snapshotValue})
}