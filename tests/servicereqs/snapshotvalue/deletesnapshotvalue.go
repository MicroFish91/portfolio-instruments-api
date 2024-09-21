package snapshotvalue

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedDeleteSnapshotValueResponse struct {
	Total float64
	Er    float64
}

func TestDeleteSnapshotValue(t *testing.T, snapId int, snapValId int, token string, expectedResponse ExpectedDeleteSnapshotValueResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/snapshots/%d/values/%d", snapId, snapValId)

	var response types.DeleteSnapshotValueResponse
	res := utils.SendDeleteRequest(t, route, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.NotEmpty(t, response.Data.Message)
		assert.EqualExportedValues(
			t,
			types.SnapshotValue{
				Snap_val_id:    snapValId,
				Snap_id:        snapId,
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

		if expectedResponse.Total != 0 {
			assert.Equal(t, expectedResponse.Total, response.Data.Snapshot_total)
		}
		if expectedResponse.Er != 0 {
			assert.Equal(t, expectedResponse.Er, response.Data.Snapshot_weighteder)
		}
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
