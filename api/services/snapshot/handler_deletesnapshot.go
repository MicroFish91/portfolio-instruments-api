package snapshot

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) DeleteSnapshot(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	snapshotParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteSnapshotParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot params from request"))
	}

	// get snapshot_values
	snapshotValues, err := h.snapshotValueStore.GetSnapshotValues(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// delete snapshot_values
	for _, sv := range snapshotValues {
		_, err := h.snapshotValueStore.DeleteSnapshotValue(c.Context(), sv.Snap_id, sv.Snap_val_id, sv.User_id)
		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}

		h.tryDeleteDeprecatedAccountById(c, sv.Account_id, sv.User_id)
		h.tryDeleteDeprecatedHoldingById(c, sv.Holding_id, sv.User_id)
	}

	// snapshot
	snapshot, err := h.snapshotStore.DeleteSnapshot(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	h.tryDeleteDeprecatedBenchmarkById(c, snapshot.Benchmark_id, snapshot.User_id)

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message":  "resource deleted successfully",
		"snapshot": snapshot,
	})
}

func (h *SnapshotHandlerImpl) tryDeleteDeprecatedAccountById(c fiber.Ctx, accountId, userId int) {
	account, err := h.accountStore.GetAccountById(c.Context(), userId, accountId)
	if err != nil {
		return
	}

	if account.Is_deprecated {
		h.accountStore.DeleteAccount(c.Context(), userId, accountId)
	}
}

func (h *SnapshotHandlerImpl) tryDeleteDeprecatedHoldingById(c fiber.Ctx, holdingId, userId int) {
	holding, err := h.holdingStore.GetHoldingById(c.Context(), userId, holdingId)
	if err != nil {
		return
	}

	if holding.Is_deprecated {
		h.holdingStore.DeleteHolding(c.Context(), userId, holdingId)
	}
}

func (h *SnapshotHandlerImpl) tryDeleteDeprecatedBenchmarkById(c fiber.Ctx, benchmarkId, userId int) {
	benchmark, err := h.benchmarkStore.GetBenchmarkById(c.Context(), userId, benchmarkId)
	if err != nil {
		return
	}

	if benchmark.Is_deprecated {
		h.benchmarkStore.DeleteBenchmark(c.Context(), userId, benchmarkId)
	}
}
