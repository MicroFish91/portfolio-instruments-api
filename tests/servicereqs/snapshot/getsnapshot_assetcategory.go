package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotByAssetCategoryResponse struct {
	HoldingNames []string
	Totals       []float64
}

func TestGetSnapshotByAssetCategory(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotByAssetCategoryResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/snapshots/%d?group_by=ASSET_CATEGORY", snapshotId)

	var getSnapshotResponse types.GetSnapshotHoldingsResponse
	res := utils.SendGetRequest(t, route, token, &getSnapshotResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, "ASSET_CATEGORY", getSnapshotResponse.Data.Field_type)

		assert.ElementsMatch(t, expectedResponse.HoldingNames, getSnapshotResponse.Data.Holdings_grouped.Fields)
		assert.ElementsMatch(t, expectedResponse.Totals, getSnapshotResponse.Data.Holdings_grouped.Total)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
