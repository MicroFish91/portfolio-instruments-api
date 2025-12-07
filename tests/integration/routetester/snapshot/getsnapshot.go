package snapshot

import (
	"fmt"
	"math"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotResponse struct {
	AccountIds    []int
	HoldingIds    []int
	Total         float64
	WeightedErPct float64
}

func TestGetSnapshot(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/snapshots/%d", snapshotId)

	var getSnapshotResponse types.GetSnapshotResponse
	res := utils.SendGetRequest(t, route, token, &getSnapshotResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)

		// Snapshot
		assert.EqualExportedValues(
			t,
			types.Snapshot{
				Snap_id:                 snapshotId,
				Description:             getSnapshotResponse.Data.Snapshot.Description,
				Snap_date:               getSnapshotResponse.Data.Snapshot.Snap_date,
				Total:                   expectedResponse.Total,
				Weighted_er_pct:         expectedResponse.WeightedErPct,
				Benchmark_id:            getSnapshotResponse.Data.Snapshot.Benchmark_id,
				Rebalance_threshold_pct: getSnapshotResponse.Data.Snapshot.Rebalance_threshold_pct,
				User_id:                 expectedUserId,
				Created_at:              getSnapshotResponse.Data.Snapshot.Created_at,
				Updated_at:              getSnapshotResponse.Data.Snapshot.Updated_at,
			},
			getSnapshotResponse.Data.Snapshot,
		)

		// Snapshot_values
		var sum float64
		for _, sv := range getSnapshotResponse.Data.Snapshot_values {
			sum += sv.Total

			assert.Equal(t, snapshotId, sv.Snap_id)
			assert.Equal(t, expectedUserId, sv.User_id)
			assert.Contains(t, expectedResponse.AccountIds, sv.Account_id)
			assert.Contains(t, expectedResponse.HoldingIds, sv.Holding_id)
		}
		assert.Equal(t, math.Round(expectedResponse.Total), math.Round(sum))

		// Accounts
		for _, a := range getSnapshotResponse.Data.Accounts {
			assert.Contains(t, expectedResponse.AccountIds, a.Account_id)
		}
		assert.Equal(t, len(expectedResponse.AccountIds), len(getSnapshotResponse.Data.Accounts))

		// Holdings
		for _, h := range getSnapshotResponse.Data.Holdings {
			assert.Contains(t, expectedResponse.HoldingIds, h.Holding_id)
		}
		assert.Equal(t, len(expectedResponse.HoldingIds), len(getSnapshotResponse.Data.Holdings))
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
