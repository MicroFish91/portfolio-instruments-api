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

type ExpectedUpdateSnapshotValueResponse struct {
	Total float64
	Er    float64
}

func TestUpdateSnapshotValue(t *testing.T, snapId int, snapValId int, payload any, token string, expectedResponse ExpectedUpdateSnapshotValueResponse, expectedUserId int, expectedStatusCode int) {
	var route string = fmt.Sprintf("/api/v1/snapshots/%d/values/%d", snapId, snapValId)

	var updateSnapshotValue types.UpdateSnapshotValueResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPut, route, token, &payload, &updateSnapshotValue)

	switch expectedStatusCode {
	case 200:
		p := payload.(snapshotvalue.UpdateSnapshotValuePayload)

		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.SnapshotValue{
				Snap_val_id:    snapValId,
				Snap_id:        snapId,
				Account_id:     p.Account_id,
				Holding_id:     p.Holding_id,
				Total:          p.Total,
				Skip_rebalance: p.Skip_rebalance,
				User_id:        expectedUserId,
				Created_at:     updateSnapshotValue.Data.Snapshot_value.Created_at,
				Updated_at:     updateSnapshotValue.Data.Snapshot_value.Updated_at,
			},
			updateSnapshotValue.Data.Snapshot_value,
		)

		if expectedResponse.Total != 0 {
			assert.Equal(t, expectedResponse.Total, updateSnapshotValue.Data.Snapshot_total)
		}
		if expectedResponse.Er != 0 {
			assert.Equal(t, expectedResponse.Er, updateSnapshotValue.Data.Snapshot_weighteder)
		}
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
