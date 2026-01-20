package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func UpdateSnapshotTestCases(t *testing.T, snapshotId int, svIds []int, benchmarkId int, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		// ---- 200 ----
		{
			Title:       "200",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-10, 0),
				Description:  "Test Description",
				Benchmark_id: benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 No Description",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-10, 0),
				Benchmark_id: benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Rebalance Threshold 1",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:               utils.Calc_target_date(-10, 0),
				Benchmark_id:            benchmarkId,
				Rebalance_threshold_pct: 0,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Rebalance Threshold 2",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:               utils.Calc_target_date(-10, 0),
				Benchmark_id:            benchmarkId,
				Rebalance_threshold_pct: 15,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Value Order",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:               utils.Calc_target_date(-10, 0),
				Benchmark_id:            benchmarkId,
				Value_order:             svIds,
				Rebalance_threshold_pct: 15,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},

		// --- 401, 404, 409 ----
		{
			Title:            "401",
			ParameterId:      snapshotId,
			ReplacementToken: tok401,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-10, 0),
				Benchmark_id: benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "404",
			ParameterId: 9999,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-12, 0),
				Benchmark_id: benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:       "409",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-1, -4),
				Benchmark_id: benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},
		{
			Title:       "409 Benchmark 1",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-10, 0),
				Benchmark_id: 0,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},
		{
			Title:       "409 Benchmark 2",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(-10, 0),
				Benchmark_id: 9999,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 ----
		{
			Title:       "400",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":    true,
				"Benchmark_id": benchmarkId,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Rebalance Threshold",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":               utils.Calc_target_date(-10, -1),
				"Benchmark_id":            benchmarkId,
				"Rebalance_threshold_pct": "15",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Value Order Type Mismatch",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":               utils.Calc_target_date(-10, -1),
				"Benchmark_id":            benchmarkId,
				"Value_order":             []string{"a", "b", "c", "d"},
				"Rebalance_threshold_pct": 10,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Value Order Ids Mismatch",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":               utils.Calc_target_date(-10, -1),
				"Benchmark_id":            benchmarkId,
				"Value_order":             []int{5, 3, 2, 1},
				"Rebalance_threshold_pct": 10,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Value Order Length Mismatch",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":               utils.Calc_target_date(-10, -1),
				"Benchmark_id":            benchmarkId,
				"Value_order":             svIds[:len(svIds)-1],
				"Rebalance_threshold_pct": 10,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Value Order No Order",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Snap_date":               utils.Calc_target_date(-10, -1),
				"Benchmark_id":            benchmarkId,
				"Value_order":             []int{},
				"Rebalance_threshold_pct": 10,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
