package snapshot

import (
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/stretchr/testify/assert"
)

type ExpectedCreateSnapshotResponse struct {
	Total         float64
	WeightedErPct float64
}

func TestCreateSnapshot(t *testing.T, payload any, token string, expectedResponse ExpectedCreateSnapshotResponse, expectedUserId int, expectedStatusCode int) (snapId int, svIds []int) {
	var createSnapshotResponse types.CreateSnapshotResponse
	res := utils.SendCreateOrUpdateRequest(t, http.MethodPost, "/api/v1/snapshots", token, &payload, &createSnapshotResponse)

	switch expectedStatusCode {
	case 201:
		p, ok := payload.(snapshot.CreateSnapshotPayload)
		if !ok {
			t.Fatal("invalid CreateSnapshotPayload")
		}

		// Verify snapshot
		assert.Equal(t, expectedStatusCode, res.StatusCode)
		assert.EqualExportedValues(
			t,
			types.Snapshot{
				Snap_id:         createSnapshotResponse.Data.Snapshot.Snap_id,
				Description:     p.Description,
				Snap_date:       p.Snap_date,
				Total:           expectedResponse.Total,
				Weighted_er_pct: expectedResponse.WeightedErPct,
				Benchmark_id:    p.Benchmark_id,
				User_id:         expectedUserId,
				Created_at:      createSnapshotResponse.Data.Snapshot.Created_at,
				Updated_at:      createSnapshotResponse.Data.Snapshot.Updated_at,
			},
			createSnapshotResponse.Data.Snapshot,
		)

		// Verify snapshot_values
		var sv_ids []int
		for i, svp := range p.Snapshot_values {
			assert.EqualExportedValues(
				t,
				types.SnapshotValue{
					Snap_val_id:    createSnapshotResponse.Data.Snapshot_values[i].Snap_val_id,
					Snap_id:        createSnapshotResponse.Data.Snapshot.Snap_id,
					Account_id:     svp.Account_id,
					Holding_id:     svp.Holding_id,
					Total:          svp.Total,
					Skip_rebalance: svp.Skip_rebalance,
					User_id:        expectedUserId,
					Created_at:     createSnapshotResponse.Data.Snapshot_values[i].Created_at,
					Updated_at:     createSnapshotResponse.Data.Snapshot_values[i].Updated_at,
				},
				createSnapshotResponse.Data.Snapshot_values[i],
			)
			sv_ids = append(sv_ids, createSnapshotResponse.Data.Snapshot_values[i].Snap_val_id)
		}

		return createSnapshotResponse.Data.Snapshot.Snap_id, svIds
	default:
		assert.Equal(t, expectedStatusCode, res.StatusCode)
	}

	return 0, nil
}
