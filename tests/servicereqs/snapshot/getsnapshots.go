package snapshot

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedGetSnapshotsResponse struct {
	Snapshots  int
	Pagination types.PaginationMetadata
}

func TestGetSnapshots(t *testing.T, route string, token string, expectedUserId int, expectedStatusCode int, expectedResponse ExpectedGetSnapshotsResponse) {
	var r string
	if route == "" {
		r = "/api/v1/snapshots"
	} else {
		r = route
	}

	var response types.GetSnapshotsResponse
	res := utils.SendGetRequest(t, r, token, &response)

	switch expectedStatusCode {
	case 200:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.Equal(t, expectedResponse.Snapshots, len(response.Data.Snapshots))
		assert.EqualExportedValues(
			t,
			expectedResponse.Pagination,
			response.Data.Pagination,
		)
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}
}
