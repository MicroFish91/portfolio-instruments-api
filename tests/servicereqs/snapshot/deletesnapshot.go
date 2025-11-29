package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSnapshot(t *testing.T, snapshotId int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/snapshots/%d", snapshotId)

	var response types.DeleteSnapshotResponse
	res := utils.SendDeleteRequest(t, route, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.NotEmpty(t, response.Data.Message)
		assert.EqualExportedValues(
			t,
			types.Snapshot{
				Snap_id:                 snapshotId,
				Description:             response.Data.Snapshot.Description,
				Snap_date:               response.Data.Snapshot.Snap_date,
				Total:                   response.Data.Snapshot.Total,
				Weighted_er_pct:         response.Data.Snapshot.Weighted_er_pct,
				Rebalance_threshold_pct: response.Data.Snapshot.Rebalance_threshold_pct,
				Benchmark_id:            response.Data.Snapshot.Benchmark_id,
				User_id:                 expectedUserId,
				Created_at:              response.Data.Snapshot.Created_at,
				Updated_at:              response.Data.Snapshot.Updated_at,
			},
			response.Data.Snapshot,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
