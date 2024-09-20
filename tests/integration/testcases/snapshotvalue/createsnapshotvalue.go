package snapshotvalue

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func CreateSnapshotValueTestCases(t *testing.T, accountIds, holdingIds []int, snapshotId, userId int, email string) []shared.TestCase {
	if len(accountIds) != 3 {
		t.Fatal("create snapshotvalue accountids incorrect length")
	}
	if len(holdingIds) != 2 {
		t.Fatal("create snapshotvalue holdingids incorrect length")
	}

	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 201 ----
		{
			Title:       "201",
			ParameterId: snapshotId,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id:     accountIds[0],
				Holding_id:     holdingIds[0],
				Total:          1750.50,
				Skip_rebalance: false,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title:       "201 Rebalance",
			ParameterId: snapshotId,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id:     accountIds[1],
				Holding_id:     holdingIds[1],
				Total:          1250.50,
				Skip_rebalance: true,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},

		{
			Title:            "401",
			ParameterId:      snapshotId,
			ReplacementToken: tok401,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id: accountIds[1],
				Holding_id: holdingIds[0],
				Total:      1000.15,
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "404",
			ParameterId: 9999,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id: accountIds[1],
				Holding_id: holdingIds[0],
				Total:      1000.15,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:       "404",
			ParameterId: snapshotId,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id: accountIds[1],
				Holding_id: 9999,
				Total:      1000.15,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:       "404",
			ParameterId: snapshotId,
			Payload: snapshotvalue.CreateSnapshotValuePayload{
				Account_id: 9999,
				Holding_id: holdingIds[0],
				Total:      1000.15,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// --- 400 ---
		{
			Title:       "400 Account Id",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Account_id": 10.1,
				"Holding_id": holdingIds[0],
				"Total":      1000.15,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Holding Id",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Account_id": accountIds[0],
				"Holding_id": true,
				"Total":      1000.15,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Total",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Account_id": accountIds[0],
				"Holding_id": holdingIds[0],
				"Total":      "One Thousand and Fifty",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
