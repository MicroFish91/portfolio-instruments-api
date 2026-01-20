package snapshot

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSnapshot(t *testing.T, snapshotId int, payload any, token string, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v2/snapshots/%d", snapshotId)

	var response types.UpdateSnapshotResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &response)

	switch expectedStatusCode {
	case 200:
		p := payload.(snapshot.UpdateSnapshotPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Snapshot{
				Snap_id:                 snapshotId,
				Description:             p.Description,
				Snap_date:               p.Snap_date,
				Total:                   response.Data.Snapshot.Total,
				Weighted_er_pct:         response.Data.Snapshot.Weighted_er_pct,
				Rebalance_threshold_pct: p.Rebalance_threshold_pct,
				Value_order:             response.Data.Snapshot.Value_order,
				Benchmark_id:            p.Benchmark_id,
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
