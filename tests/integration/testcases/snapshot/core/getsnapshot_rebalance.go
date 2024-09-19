package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotRebalanceTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotRebalanceResponse{
			Target_allocation: []types.AssetAllocation{
				{
					Category: "TSM",
					Value:    1500.90,
				},
				{
					Category: "ITB",
					Value:    1000.60,
				},
			},
			Current_allocation: []types.AssetAllocation{
				{
					Category: "TSM",
					Value:    1001.00,
				},
				{
					Category: "ITB",
					Value:    1500.50,
				},
			},
			Change_required: []types.AssetAllocation{
				{
					Category: "TSM",
					Value:    499.90,
				},
				{
					Category: "ITB",
					Value:    -499.90,
				},
			},
			Snapshot_total:            2501.5,
			Snapshot_total_omit_skips: 2501.5,
			Rebalance_thresh_pct:      50,
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
