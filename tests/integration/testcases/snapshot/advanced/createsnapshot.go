package snapshot

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetCreateSnapshotAdvancedTestCase(t *testing.T, benchmarkId int, accountIds []int, holdingIds []int) shared.PostTestCase {
	if len(accountIds) != 9 {
		t.Fatal("unexpected accountId length for creating advanced snapshot")
	}
	if len(holdingIds) != 14 {
		t.Fatal("unexpeted holdingId length for creating advanced snapshot")
	}

	return shared.PostTestCase{
		Title: "200",
		Payload: snapshot.CreateSnapshotPayload{
			Snap_date: utils.Calc_target_date(0, -3),
			Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
				// Vanguard - Taxable
				// Total (excluding skip): 32801.48
				// Total (including skip): 36427.58
				{Account_id: accountIds[0], Holding_id: holdingIds[2], Total: 10341.01, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[12], Total: 11979.70, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[6], Total: 1092.52, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[3], Total: 3683.72, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[7], Total: 1913.09, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[5], Total: 3891.44, Skip_rebalance: false},
				{Account_id: accountIds[0], Holding_id: holdingIds[9], Total: 3626.10, Skip_rebalance: true},

				// Vanguard - Roth
				// Total (excluding skip): 51919.82
				// Total (including skip): 63578.53
				{Account_id: accountIds[1], Holding_id: holdingIds[3], Total: 9371.45, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[1], Total: 2271.85, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[13], Total: 12096.14, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[10], Total: 7020.39, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[8], Total: 13431.37, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[4], Total: 7728.62, Skip_rebalance: false},
				{Account_id: accountIds[1], Holding_id: holdingIds[0], Total: 11658.71, Skip_rebalance: true},

				// Vanguard - Traditional
				// Total (excluding skip): 53729.67
				// Total (including skip): 59610.47
				{Account_id: accountIds[2], Holding_id: holdingIds[4], Total: 7780.79, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[9], Total: 6969.85, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[11], Total: 10443.17, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[5], Total: 12502.99, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[2], Total: 286.65, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[3], Total: 15706.22, Skip_rebalance: false},
				{Account_id: accountIds[2], Holding_id: holdingIds[6], Total: 5880.80, Skip_rebalance: true},

				// Fidelity - Taxable
				// Total (excluding skip): 36904.09
				// Total (including skip): 37341.87
				{Account_id: accountIds[3], Holding_id: holdingIds[3], Total: 7625.72, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[7], Total: 10586.65, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[4], Total: 7324.53, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[9], Total: 1168.83, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[0], Total: 5268.47, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[2], Total: 4929.89, Skip_rebalance: false},
				{Account_id: accountIds[3], Holding_id: holdingIds[11], Total: 437.78, Skip_rebalance: true},

				// Fidelity - Roth
				// Total (excluding skip): 45789.25
				// Total (including skip): 59484.99
				{Account_id: accountIds[4], Holding_id: holdingIds[8], Total: 15407.15, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[9], Total: 4662.95, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[4], Total: 13466.88, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[12], Total: 4747.47, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[10], Total: 515.56, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[6], Total: 6889.24, Skip_rebalance: false},
				{Account_id: accountIds[4], Holding_id: holdingIds[4], Total: 13695.74, Skip_rebalance: true},

				// Fidelity - Traditional
				// Total (excluding skip): 34876.01
				// Total (including skip): 44324.53
				{Account_id: accountIds[5], Holding_id: holdingIds[4], Total: 12673.13, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[12], Total: 3073.30, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[2], Total: 11547.25, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[6], Total: 2240.44, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[1], Total: 3483.74, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[7], Total: 1858.15, Skip_rebalance: false},
				{Account_id: accountIds[5], Holding_id: holdingIds[13], Total: 9448.52, Skip_rebalance: true},

				// Schwab - Taxable
				// Total (excluding skip): 60167.70
				// Total (including skip): 71022.17
				{Account_id: accountIds[6], Holding_id: holdingIds[1], Total: 15130.08, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[13], Total: 11453.51, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[11], Total: 7429.49, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[3], Total: 1639.83, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[5], Total: 14835.84, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[12], Total: 9678.95, Skip_rebalance: false},
				{Account_id: accountIds[6], Holding_id: holdingIds[10], Total: 10854.47, Skip_rebalance: true},

				// Schwab - Roth
				// Total (excluding skip): 61823.27
				// Total (including skip): 66100.27
				{Account_id: accountIds[7], Holding_id: holdingIds[6], Total: 15737.92, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[12], Total: 9434.05, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[5], Total: 10096.92, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[13], Total: 15411.10, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[2], Total: 6899.97, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[11], Total: 4343.31, Skip_rebalance: false},
				{Account_id: accountIds[7], Holding_id: holdingIds[4], Total: 4277.00, Skip_rebalance: true},

				// Schwab - Traditional
				// Total (excluding skip): 57298.58
				// Total (including skip): 69157.69
				{Account_id: accountIds[8], Holding_id: holdingIds[2], Total: 14045.3, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[9], Total: 11642.63, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[6], Total: 6618.72, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[8], Total: 1124.46, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[5], Total: 13441.36, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[0], Total: 10426.11, Skip_rebalance: false},
				{Account_id: accountIds[8], Holding_id: holdingIds[13], Total: 11859.11, Skip_rebalance: true},
			},
			Benchmark_id: benchmarkId,
		},
		ExpectedResponse: snapshotTester.ExpectedCreateSnapshotResponse{
			Total:         507108.10,
			WeightedErPct: 0.389,
		},
		ExpectedStatusCode: fiber.StatusCreated,
	}
}
