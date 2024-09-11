package snapshot

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotByMaturationDateResponse struct {
	Resources        []types.MaturationDateResource
	Maturation_start string
	Maturation_end   string
}

func TestGetSnapshotByMaturationDate(t *testing.T, snapshotId int, token string, expectedResponse ExpectedGetSnapshotByMaturationDateResponse, expectedUserId int, expectedStatusCode int) {
	var route = fmt.Sprintf("/api/v1/snapshots/%d?group_by=MATURATION_DATE", snapshotId)
	if expectedResponse.Maturation_start != "" {
		route = fmt.Sprintf("%s&maturation_start=%s", route, expectedResponse.Maturation_start)
	}
	if expectedResponse.Maturation_end != "" {
		route = fmt.Sprintf("%s&maturation_end=%s", route, expectedResponse.Maturation_end)
	}

	var getSnapshotResponse types.GetSnapshotMaturationDateResponse
	res := utils.SendGetRequest(t, route, token, &getSnapshotResponse)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, "MATURATION_DATE", getSnapshotResponse.Data.Field_type)

		assert.Equal(t, expectedResponse.Maturation_start, getSnapshotResponse.Data.Maturation_start)
		assert.Equal(t, expectedResponse.Maturation_end, getSnapshotResponse.Data.Maturation_end)

		assert.Equal(t, len(expectedResponse.Resources), len(getSnapshotResponse.Data.Resources))

		for _, resource := range getSnapshotResponse.Data.Resources {
			assertHasMatchingResource(t, resource, expectedResponse.Resources)
		}
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}

func assertHasMatchingResource(t *testing.T, resource types.MaturationDateResource, expectedResources []types.MaturationDateResource) {
	hasMatch := false
	for _, er := range expectedResources {
		if resource.Account_name == er.Account_name && resource.Holding_name == er.Holding_name {
			expectedResource := types.MaturationDateResource{
				Account_name:      er.Account_name,
				Holding_name:      er.Holding_name,
				Asset_category:    er.Asset_category,
				Interest_rate_pct: resource.Interest_rate_pct,
				Maturation_date:   resource.Maturation_date,
				Total:             er.Total,
				Skip_rebalance:    er.Skip_rebalance,
			}

			assert.EqualExportedValues(
				t,
				expectedResource,
				resource,
			)

			hasMatch = true
		}
	}

	if !hasMatch {
		assert.Fail(t, "no matching get snapshot maturation date resource")
	}
}
