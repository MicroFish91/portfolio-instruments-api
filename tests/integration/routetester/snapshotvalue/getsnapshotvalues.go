package snapshotvalue

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotValuesResponse struct {
	SnapshotValues int
}

func TestGetSnapshotValues(t *testing.T, token string, snapshotId, expectedUserId, expectedStatusCode int, expectedResponse ExpectedGetSnapshotValuesResponse) {
	var route = fmt.Sprintf("/api/v2/snapshots/%d/values", snapshotId)

	var response types.GetSnapshotValuesResponse
	res := utils.SendGetRequest(t, route, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, expectedResponse.SnapshotValues, len(response.Data.Snapshot_values))
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
