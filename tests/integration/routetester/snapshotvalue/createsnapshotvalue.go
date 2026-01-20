package snapshotvalue

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateSnapshotValue(t *testing.T, payload any, token string, sid, expectedUserId, expectedStatusCode int) int {
	var response types.CreateSnapshotValueResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPost, fmt.Sprintf("/api/v2/snapshots/%d/values", sid), token, &payload, &response)

	switch expectedStatusCode {
	case 201:
		p := payload.(snapshotvalue.CreateSnapshotValuePayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.SnapshotValue{
				Snap_val_id:    response.Data.Snapshot_value.Snap_val_id,
				Snap_id:        sid,
				Account_id:     p.Account_id,
				Holding_id:     p.Holding_id,
				Total:          p.Total,
				Skip_rebalance: p.Skip_rebalance,
				User_id:        expectedUserId,
				Created_at:     response.Data.Snapshot_value.Created_at,
				Updated_at:     response.Data.Snapshot_value.Updated_at,
			},
			response.Data.Snapshot_value,
		)
		return response.Data.Snapshot_value.Snap_val_id
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return response.Data.Snapshot_value.Snap_val_id
}
