package snapshot

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/gofiber/fiber/v3"
)

func GetCreateSnapshotTestCases(t *testing.T, userId int, email string) []shared.PostTestCase {
	return []shared.PostTestCase{
		// ---- 200 ----
		{
			Title: "200",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date:   "",
				Description: "",
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{
						Account_id:     0,
						Holding_id:     0,
						Total:          0,
						Skip_rebalance: false,
					},
				},
				Benchmark_id: 0,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
	}
}
