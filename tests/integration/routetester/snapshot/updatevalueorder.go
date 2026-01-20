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

func TestUpdateSnapshotValueOrder(t *testing.T, snapshotId int, payload any, token string, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v2/snapshots/%d/order", snapshotId)

	var response types.UpdateValueOrderResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &response)

	switch expectedStatusCode {
	case 200:
		p := payload.(snapshot.UpdateValueOrderPayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Snapshot{
				Snap_id:                 snapshotId,
				Description:             response.Data.Snapshot.Description,
				Snap_date:               response.Data.Snapshot.Snap_date,
				Total:                   response.Data.Snapshot.Total,
				Weighted_er_pct:         response.Data.Snapshot.Weighted_er_pct,
				Rebalance_threshold_pct: response.Data.Snapshot.Rebalance_threshold_pct,
				Value_order:             p.Value_order,
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
