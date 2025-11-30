package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotRebalanceTestCase(*testing.T) shared.TestCase {
	return shared.TestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotRebalanceResponse{
			Target_allocation: []types.AssetAllocation{
				{
					Category: "TSM",
					Value:    87073.97,
				},
				{
					Category: "LTB",
					Value:    87073.97,
				},
				{
					Category: "CASH",
					Value:    87073.97,
				},
				{
					Category: "GOLD",
					Value:    87073.97,
				},
				{
					Category: "DSCV",
					Value:    87073.97,
				},
			},
			Current_allocation: []types.AssetAllocation{
				{
					Category: "CASH",
					Value:    72494.33,
				},
				{
					Category: "OTHER",
					Value:    93682.02,
				},
				{
					Category: "STB",
					Value:    32578.84,
				},
				{
					Category: "GOLD",
					Value:    45562.89,
				},
				{
					Category: "TSM",
					Value:    30052.47,
				},
				{
					Category: "LTB",
					Value:    50848.65,
				},
				{
					Category: "ITB",
					Value:    38960.75,
				},
				{
					Category: "DSCV",
					Value:    71189.92,
				},
			},
			Change_required: []types.AssetAllocation{
				{
					Category: "TSM",
					Value:    57021.50,
				},
				{
					Category: "LTB",
					Value:    36225.32,
				},
				{
					Category: "CASH",
					Value:    14579.64,
				},
				{
					Category: "GOLD",
					Value:    41511.08,
				},
				{
					Category: "DSCV",
					Value:    15884.05,
				},
				{
					Category: "OTHER",
					Value:    -93682.02,
				},
				{
					Category: "STB",
					Value:    -32578.84,
				},
				{
					Category: "ITB",
					Value:    -38960.75,
				},
			},
			Snapshot_total:            507108.10,
			Snapshot_total_omit_skips: 435369.87,
			Rebalance_deviation_pct:   13,
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
