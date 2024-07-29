package snapshot

import (
	"errors"
	"math"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *SnapshotHandlerImpl) RebalanceSnapshot(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user request body"))
	}

	snapshotParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(GetSnapshotByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid snapshot params from request"))
	}

	// snapshot
	snapshot, _, err := h.snapshotStore.GetSnapshotById(c.Context(), snapshotParams.Id, userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	if snapshot.Benchmark_id == 0 {
		return utils.SendError(c, fiber.StatusConflict, errors.New("snapshot must include a target benchmark_id for a rebalance computation to take place"))
	}

	// benchmark
	benchmark, err := h.benchmarkStore.GetBenchmarkById(c.Context(), userPayload.User_id, snapshot.Benchmark_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// holdings
	holdingsGrouped, err := h.snapshotStore.TallyByHolding(c.Context(), userPayload.User_id, snapshot.Snap_id, &types.GetTallyByHoldingStoreOptions{
		Tally_by:      types.BY_ASSET_CATEGORY,
		Omit_skip_reb: true,
	})
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	// compute rebalance
	target, current, change, rebThreshPct, err := h.computeRebalance(benchmark.Asset_allocation, *holdingsGrouped, snapshot.Total)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"target_allocation":    target,
		"current_allocation":   current,
		"change_required":      change,
		"target_deviation_pct": rebThreshPct,
	})
}

type AssetAllocation struct {
	Category string  `json:"category"`
	Value    float64 `json:"value"`
}

func (h *SnapshotHandlerImpl) computeRebalance(balloc []types.AssetAllocationPct, halloc types.ResourcesGrouped, total float64) (target *[]AssetAllocation, current *[]AssetAllocation, change *[]AssetAllocation, rebThreshPct int, e error) {
	if len(halloc.Fields) != len(halloc.Total) {
		return nil, nil, nil, 0, errors.New("internal: could not compute rebalance, mismatching number of fields and totals for grouped holdings")
	}

	// Benchmark target allocation
	var tar []AssetAllocation
	for _, b := range balloc {
		var t AssetAllocation
		t.Category = string(b.Category)
		t.Value = math.Round(float64(b.Percent)/float64(100)*total*100) / 100 // Potential for small rounding errors, however, decimal point accuracy for this value is not super important
		tar = append(tar, t)
	}

	// Current holding allocation
	var cur []AssetAllocation
	for i := 0; i < len(halloc.Fields); i++ {
		var c AssetAllocation
		c.Category = halloc.Fields[i]
		c.Value = halloc.Total[i]
		cur = append(cur, c)
	}

	// Rebalance diff required
	ch, rebThresh, err := h.computeRebalanceDiff(tar, cur, total)
	if err != nil {
		return nil, nil, nil, 0, err
	}

	return &tar, &cur, ch, rebThresh, nil
}

func (h *SnapshotHandlerImpl) computeRebalanceDiff(target []AssetAllocation, current []AssetAllocation, total float64) (alloc *[]AssetAllocation, rebThreshPct int, e error) {
	var (
		maxDeviation = 0
		chmap        = make(map[string]float64)
	)

	// Compute diffs for the target (benchmark) asset categories
	for _, t := range target {
		// If we've already assigned a value for the category, then skip
		if _, ok := chmap[t.Category]; ok {
			continue
		}

		// Check if this value also exists in our current holding allocation; use it to compute the diff
		var alloc AssetAllocation
		for _, a := range current {
			if a.Category == t.Category {
				alloc = a
				break
			}
		}

		diff := t.Value - alloc.Value
		chmap[t.Category] = diff

		deviation := math.Round(diff / total * 100)
		if int(math.Abs(deviation)) > maxDeviation {
			maxDeviation = int(math.Abs(deviation))
		}
	}

	// Compute any remaining diffs for the non-target (non-benchmark) asset categories
	for _, c := range current {
		// If we've already assigned a value for the category, then skip
		if _, ok := chmap[c.Category]; ok {
			continue
		}
		chmap[c.Category] = -c.Value
	}

	// Transform map into an allocation slice
	var ch []AssetAllocation
	for key, value := range chmap {
		var alloc AssetAllocation
		alloc.Category = key
		alloc.Value = value
		ch = append(ch, alloc)
	}

	return &ch, maxDeviation, nil
}
