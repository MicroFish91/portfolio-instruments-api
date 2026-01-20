package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func UpdateSnapshotValueOrderTestCases(t *testing.T, snapshotId int, svIds []int, benchmarkId int, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	// Set a new order for the values, swapping the first and last element
	newValueOrder := make([]int, len(svIds))
	copy(newValueOrder, svIds)
	newValueOrder[0], newValueOrder[len(newValueOrder)-1] = svIds[len(svIds)-1], svIds[0]

	duplicateValueOrder := make([]int, len(svIds))
	copy(duplicateValueOrder, svIds)
	duplicateValueOrder[len(duplicateValueOrder)-1] = duplicateValueOrder[0]

	return []testcases.TestCase{
		// ---- 200 ----
		{
			Title:       "200 Swapped Order",
			ParameterId: snapshotId,
			Payload: snapshot.UpdateValueOrderPayload{
				Value_order: newValueOrder,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},

		// --- 401, 404, 409 ----
		{
			Title:            "401",
			ParameterId:      snapshotId,
			ReplacementToken: tok401,
			Payload: snapshot.UpdateSnapshotPayload{
				Value_order: svIds,
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "404",
			ParameterId: 9999,
			Payload: snapshot.UpdateSnapshotPayload{
				Value_order: svIds,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// ---- 400 ----
		{
			Title:       "400 No Order",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Value_order": []int{},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Type Mismatch",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Value_order": []any{"a", "b", "c", "d"},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Length Mismatch",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Value_order": svIds[:len(svIds)-1],
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Extra Duplicate",
			ParameterId: snapshotId,
			Payload: map[string]any{
				"Value_order": duplicateValueOrder,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
