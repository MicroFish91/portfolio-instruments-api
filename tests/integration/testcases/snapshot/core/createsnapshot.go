package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetCreateSnapshotTestCases(t *testing.T, benchmarkId int, accountIds, holdingIds []int, userId int, email string) []shared.PostTestCase {
	if len(accountIds) != 3 {
		t.Fatal("unexpected accountId length for creating core snapshot")
	}
	if len(holdingIds) != 2 {
		t.Fatal("unexpected holdingId length for creating core snapshot")
	}

	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.PostTestCase{
		// ---- 201 ----
		{
			Title: "201",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
					{Account_id: accountIds[0], Holding_id: holdingIds[1], Total: 500.50, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[0], Total: 750.75, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[1], Total: 1000.00, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse: snapshotTester.ExpectedCreateSnapshotResponse{
				Total:         2501.50,
				WeightedErPct: 0.180,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 No Benchmark",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
					{Account_id: accountIds[0], Holding_id: holdingIds[1], Total: 500.50, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[0], Total: 750.75, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[1], Total: 1000.00, Skip_rebalance: false},
				},
			},
			ExpectedResponse: snapshotTester.ExpectedCreateSnapshotResponse{
				Total:         2501.50,
				WeightedErPct: 0.180,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Description",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date:   utils.Calc_target_date(0, -3),
				Description: "With Description",
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
					{Account_id: accountIds[0], Holding_id: holdingIds[1], Total: 500.50, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[0], Total: 750.75, Skip_rebalance: false},
					{Account_id: accountIds[1], Holding_id: holdingIds[1], Total: 1000.00, Skip_rebalance: false},
				},
			},
			ExpectedResponse: snapshotTester.ExpectedCreateSnapshotResponse{
				Total:         2501.50,
				WeightedErPct: 0.180,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},

		// ---- 401 ----
		{
			Title:            "401",
			ReplacementToken: tok401,
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},

		// ---- 409 ----
		{
			Title: "409 Benchmark",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: 9999,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusConflict,
		},
		{
			Title: "409 Account",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: 9999, Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusConflict,
		},
		{
			Title: "409 Holding",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: 9999, Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 -----
		{
			Title: "400 No Date",
			Payload: snapshot.CreateSnapshotPayload{
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Future Date",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, 3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Empty Snapshot",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date:    utils.Calc_target_date(0, -3),
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Date 1",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: "Oct 23, 2010",
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Date 2",
			Payload: map[string]any{
				"Snap_date": 10,
				"Snapshot_values": []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				"Benchmark_id": benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Account",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Holding_id: holdingIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Account",
			Payload: map[string]any{
				"Snap_date": utils.Calc_target_date(0, -3),
				"Snapshot_values": []map[string]any{
					{"Account_id": "1", "Holding_id": holdingIds[0], "Total": 250.25, "Skip_rebalance": false},
				},
				"Benchmark_id": benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Holding",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Total: 250.25, Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Holding",
			Payload: map[string]any{
				"Snap_date": utils.Calc_target_date(0, -3),
				"Snapshot_values": []map[string]any{
					{"Account_id": accountIds[0], "Holding_id": true, "Total": 250.25, "Skip_rebalance": false},
				},
				"Benchmark_id": benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Total",
			Payload: snapshot.CreateSnapshotPayload{
				Snap_date: utils.Calc_target_date(0, -3),
				Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Skip_rebalance: false},
				},
				Benchmark_id: benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Total",
			Payload: map[string]any{
				"Snap_date": utils.Calc_target_date(0, -3),
				"Snapshot_values": []map[string]any{
					{"Account_id": accountIds[0], "Holding_id": holdingIds[0], "Total": "abcd", "Skip_rebalance": false},
				},
				"Benchmark_id": benchmarkId,
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Benchmark",
			Payload: map[string]any{
				"Snap_date": utils.Calc_target_date(0, -3),
				"Snapshot_values": []snapshotvalue.CreateSnapshotValuePayload{
					{Account_id: accountIds[0], Holding_id: holdingIds[0], Skip_rebalance: false},
				},
				"Benchmark_id": []int{1, 2, 3},
			},
			ExpectedResponse:   snapshotTester.ExpectedCreateSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
