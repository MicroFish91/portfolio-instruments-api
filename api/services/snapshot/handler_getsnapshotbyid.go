package snapshot

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) GetSnapshotById(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	queryPayload, ok := c.Locals(constants.LOCALS_REQ_QUERY).(GetSnapshotByIdQuery)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid query params from request"))
	}

	snapshotParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(GetSnapshotByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot params from request"))
	}

	// Todo: finish adding cases to cover the various tally scenarios
	if queryPayload.Tally_by != "" {
		return h.tallyBySnapshotHandler(c, queryPayload.Tally_by, snapshotParams.Id, userPayload.User_id)
	}

	// snapshot, snapshotValues
	snapshot, snapshotValues, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	accountIds, holdingIds := h.gatherSnapshotResourceIds(snapshotValues)

	// accounts
	accounts, err := h.getSnapshotAccounts(c, accountIds, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// holdings
	holdings, err := h.getSnapshotHoldings(c, holdingIds, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"snapshot":        snapshot,
		"snapshot_values": snapshotValues,
		"accounts":        accounts,
		"holdings":        holdings,
	})
}

func (h *SnapshotHandlerImpl) getSnapshotAccounts(c fiber.Ctx, accountsIds *[]int, userId int) (*[]types.Account, error) {
	accounts, _, err := h.accountStore.GetAccounts(
		c.Context(),
		userId,
		&types.GetAccountsStoreOptions{
			AccountIds: *accountsIds,
		},
	)

	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (h *SnapshotHandlerImpl) getSnapshotHoldings(c fiber.Ctx, holdingIds *[]int, userId int) (*[]types.Holding, error) {
	holdings, _, err := h.holdingStore.GetHoldings(
		c.Context(),
		userId,
		&types.GetHoldingsStoreOptions{
			Holding_ids: *holdingIds,
			Page_size:   100,
		},
	)

	if err != nil {
		return nil, err
	}
	return holdings, nil
}

func (h *SnapshotHandlerImpl) gatherSnapshotResourceIds(snapshotValues *[]types.SnapshotValues) (accIds *[]int, holdIds *[]int) {
	accIdsSet := map[int]struct{}{}
	holdIdsSet := map[int]struct{}{}

	for _, sv := range *snapshotValues {
		accIdsSet[sv.Account_id] = struct{}{}
		holdIdsSet[sv.Holding_id] = struct{}{}
	}

	accountIds := make([]int, 0, len(accIdsSet))
	for key := range accIdsSet {
		accountIds = append(accountIds, key)
	}

	holdingIds := make([]int, 0, len(holdIdsSet))
	for key := range holdIdsSet {
		holdingIds = append(holdingIds, key)
	}

	return &accountIds, &holdingIds
}

func (h *SnapshotHandlerImpl) tallyBySnapshotHandler(c fiber.Ctx, tc TallyCategory, snapId, userId int) error {
	switch tc {
	case BY_ACCOUNT:
		// Todo: Add store method
		return nil
	case BY_TAX_SHELTER:
		// Todo: Add store method
		return nil
	case BY_ASSET_CATEGORY:
		// Todo: Add store method
		return nil
	default:
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("provided an unsupported tally_by category"))
	}
}
