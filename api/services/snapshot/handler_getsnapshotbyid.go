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

	// snapshot, snapshotValues
	snapshot, snapshotValues, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// group_by
	if queryPayload.Group_by != "" {
		return h.handleGroupByResource(c, snapshotParams.Id, userPayload.User_id, HandleGroupByResourceOptions(queryPayload))
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

func (h *SnapshotHandlerImpl) getSnapshotAccounts(c fiber.Ctx, accountsIds []int, userId int) ([]types.Account, error) {
	accounts, _, err := h.accountStore.GetAccounts(
		c.Context(),
		userId,
		&types.GetAccountsStoreOptions{
			AccountIds: accountsIds,
		},
	)

	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (h *SnapshotHandlerImpl) getSnapshotHoldings(c fiber.Ctx, holdingIds []int, userId int) ([]types.Holding, error) {
	holdings, _, err := h.holdingStore.GetHoldings(
		c.Context(),
		userId,
		&types.GetHoldingsStoreOptions{
			Holding_ids: holdingIds,
			Page_size:   100,
		},
	)

	if err != nil {
		return nil, err
	}
	return holdings, nil
}

// Takes a complete slice of snapshot_values and returns all unique account and holding ids
func (h *SnapshotHandlerImpl) gatherSnapshotResourceIds(snapshotValues []types.SnapshotValue) (accIds []int, holdIds []int) {
	accIdsSet := map[int]struct{}{}
	holdIdsSet := map[int]struct{}{}

	for _, sv := range snapshotValues {
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

	return accountIds, holdingIds
}

type HandleGroupByResourceOptions struct {
	Group_by         string
	Maturation_start string
	Maturation_end   string
}

// Handles all logic paths for resolving the tally_by query parameter
func (h *SnapshotHandlerImpl) handleGroupByResource(c fiber.Ctx, snapId, userId int, options HandleGroupByResourceOptions) error {
	switch GroupByCategory(options.Group_by) {
	case BY_ACCOUNT_NAME, BY_ACCOUNT_INSTITUTION, BY_TAX_SHELTER:

		var groupBy types.AccountsGroupByCategory
		if GroupByCategory(options.Group_by) == BY_ACCOUNT_NAME {
			groupBy = types.BY_ACCOUNT_NAME
		} else if GroupByCategory(options.Group_by) == BY_ACCOUNT_INSTITUTION {
			groupBy = types.BY_ACCOUNT_INSTITUTION
		} else {
			groupBy = types.BY_TAX_SHELTER
		}

		accountsGrouped, err := h.snapshotStore.GroupByAccount(c.Context(), snapId, userId, &types.GetGroupByAccountStoreOptions{
			Group_by: groupBy,
		})

		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}
		return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
			"accounts_grouped": accountsGrouped,
			"field_type":       GroupByCategory(groupBy),
		})

	case BY_ASSET_CATEGORY:

		var groupBy types.HoldingsGroupByCategory
		if GroupByCategory(options.Group_by) == BY_ASSET_CATEGORY {
			groupBy = types.BY_ASSET_CATEGORY
		}

		holdingsGrouped, err := h.snapshotStore.GroupByHolding(c.Context(), userId, snapId, &types.GetGroupByHoldingStoreOptions{
			Group_by:      groupBy,
			Omit_skip_reb: false,
		})

		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}
		return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
			"holdings_grouped": holdingsGrouped,
			"field_type":       GroupByCategory(options.Group_by),
		})

	case BY_MATURATION_DATE:

		resources, err := h.snapshotStore.GroupByMaturationDate(c.Context(), userId, snapId, &types.GetGroupByMaturationDateStoreOptions{
			Maturation_start: options.Maturation_start,
			Maturation_end:   options.Maturation_end,
		})

		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}
		return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
			"resources":        resources,
			"field_type":       GroupByCategory(options.Group_by),
			"maturation_start": options.Maturation_start,
			"maturation_end":   options.Maturation_end,
		})

	case BY_LIQUIDITY:

		resources, liquidTotal, err := h.snapshotStore.GroupByLiquidity(c.Context(), userId, snapId)
		if err != nil {
			return utils.SendError(c, utils.StatusCodeFromError(err), err)
		}

		return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
			"resources":    resources,
			"liquid_total": liquidTotal,
			"field_type":   GroupByCategory(options.Group_by),
		})

	default:
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("provided an unsupported group_by request category"))
	}
}
