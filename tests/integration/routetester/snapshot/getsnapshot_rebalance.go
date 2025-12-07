package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotRebalanceResponse struct {
	Target_allocation         []types.AssetAllocation
	Current_allocation        []types.AssetAllocation
	Change_required           []types.AssetAllocation
	Snapshot_total            float64
	Snapshot_total_omit_skips float64
	Rebalance_deviation_pct   int
	Needs_rebalance           bool
}

func TestGetSnapshotRebalance(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotRebalanceResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/snapshots/%d/rebalance", snapshotId)

	var response types.GetSnapshotRebalanceResponse
	res := utils.SendGetRequest(t, route, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.ElementsMatch(t, expectedResponse.Target_allocation, *response.Data.Target_allocation)
		assert.ElementsMatch(t, expectedResponse.Current_allocation, *response.Data.Current_allocation)
		assert.ElementsMatch(t, expectedResponse.Change_required, *response.Data.Change_required)
		assert.Equal(t, expectedResponse.Snapshot_total, response.Data.Snapshot_total)
		assert.Equal(t, expectedResponse.Snapshot_total_omit_skips, response.Data.Snapshot_total_omit_skips)
		assert.Equal(t, expectedResponse.Rebalance_deviation_pct, response.Data.Rebalance_deviation_pct)
		assert.Equal(t, expectedResponse.Needs_rebalance, response.Data.Needs_rebalance)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
