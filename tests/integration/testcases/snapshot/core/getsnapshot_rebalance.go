package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotRebalanceTestCases(t *testing.T, snapWithBenchmarkId, snapWithBenchmarkNoThreshId, snapWithoutBenchmarkId, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:       "200",
			ParameterId: snapWithBenchmarkId,
			ExpectedResponse: routeTester.ExpectedGetSnapshotRebalanceResponse{
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
				Rebalance_deviation_pct:   20,
				Needs_rebalance:           true,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200",
			ParameterId: snapWithBenchmarkNoThreshId,
			ExpectedResponse: routeTester.ExpectedGetSnapshotRebalanceResponse{
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
				Rebalance_deviation_pct:   20,
				Needs_rebalance:           false,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        snapWithBenchmarkId,
			ReplacementToken:   tok401,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        9999,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:              "409",
			ParameterId:        snapWithoutBenchmarkId,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusConflict,
		},
	}
}
