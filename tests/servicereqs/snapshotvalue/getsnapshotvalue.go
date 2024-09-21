package snapshotvalue

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetSnapshotValue(t *testing.T, sid, svid int, token string, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/snapshots/%d/values/%d", sid, svid)

	var response types.GetSnapshotValueResponse
	res := utils.SendGetRequest(t, route, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.SnapshotValue{
				Snap_val_id:    svid,
				Snap_id:        sid,
				Account_id:     response.Data.Snapshot_value.Account_id,
				Holding_id:     response.Data.Snapshot_value.Holding_id,
				Total:          response.Data.Snapshot_value.Total,
				Skip_rebalance: response.Data.Snapshot_value.Skip_rebalance,
				User_id:        expectedUserId,
				Created_at:     response.Data.Snapshot_value.Created_at,
				Updated_at:     response.Data.Snapshot_value.Updated_at,
			},
			response.Data.Snapshot_value,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
