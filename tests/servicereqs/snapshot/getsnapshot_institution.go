package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotByInstitutionResponse struct {
	Institutions []string
	Totals       []float64
}

func TestGetSnapshotByInstitution(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotByInstitutionResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v2/snapshots/%d?group_by=ACCOUNT_INSTITUTION", snapshotId)

	var getSnapshotResponse types.GetSnapshotAccountsResponse
	res := utils.SendGetRequest(t, route, token, &getSnapshotResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, "ACCOUNT_INSTITUTION", getSnapshotResponse.Data.Field_type)

		assert.ElementsMatch(t, expectedResponse.Institutions, getSnapshotResponse.Data.Accounts_grouped.Fields)
		assert.ElementsMatch(t, expectedResponse.Totals, getSnapshotResponse.Data.Accounts_grouped.Total)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
