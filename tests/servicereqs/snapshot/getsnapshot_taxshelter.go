package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotByTaxShelterResponse struct {
	TaxShelters []string
	Totals      []float64
}

func TestGetSnapshotByTaxShelter(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotByTaxShelterResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/snapshots/%d?group_by=TAX_SHELTER", snapshotId)

	var getSnapshotResponse types.GetSnapshotAccountsResponse
	res := utils.SendGetRequest(t, route, token, &getSnapshotResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, "TAX_SHELTER", getSnapshotResponse.Data.Field_type)

		assert.ElementsMatch(t, expectedResponse.TaxShelters, getSnapshotResponse.Data.Accounts_grouped.Fields)
		assert.ElementsMatch(t, expectedResponse.Totals, getSnapshotResponse.Data.Accounts_grouped.Total)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}