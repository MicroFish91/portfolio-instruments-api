package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotRebalanceTestCases(t *testing.T, snapWithBenchmarkId, snapWithoutBenchmarkId, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		{
			Title:       "200",
			ParameterId: snapWithBenchmarkId,
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
		},
		{
			Title:              "401",
			ParameterId:        snapWithBenchmarkId,
			ReplacementToken:   tok401,
			ExpectedResponse:   snapshotTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        9999,
			ExpectedResponse:   snapshotTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:              "409",
			ParameterId:        snapWithoutBenchmarkId,
			ExpectedResponse:   snapshotTester.ExpectedGetSnapshotRebalanceResponse{},
			ExpectedStatusCode: fiber.StatusConflict,
		},
	}
}
