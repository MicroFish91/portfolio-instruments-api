package snapshotvalue

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func UpdateSnapshotValueTestCases(t *testing.T, accountIds, holdingIds []int, sid, svid, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:        "200",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: accountIds[0],
				Holding_id: holdingIds[0],
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:            "401",
			ParameterId:      sid,
			ParameterId2:     svid,
			ReplacementToken: tok401,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: accountIds[0],
				Holding_id: holdingIds[0],
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},

		// ---- 404 ----
		{
			Title:        "404 sid",
			ParameterId:  99999,
			ParameterId2: svid,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: accountIds[0],
				Holding_id: holdingIds[0],
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:        "404 svid",
			ParameterId:  sid,
			ParameterId2: 99999,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: accountIds[0],
				Holding_id: holdingIds[0],
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:        "404 Account",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: 99999,
				Holding_id: holdingIds[0],
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:        "404 Holding",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: accountIds[0],
				Holding_id: 99999,
				Total:      1337.25,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// ---- 400 ----
		{
			Title:        "400 Account Id",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: map[string]any{
				"Account_id": 10.1,
				"Holding_id": holdingIds[0],
				"Total":      1000.15,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:        "400 Holding Id",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: map[string]any{
				"Account_id": accountIds[0],
				"Holding_id": true,
				"Total":      1000.15,
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:        "400 Total",
			ParameterId:  sid,
			ParameterId2: svid,
			Payload: map[string]any{
				"Account_id": accountIds[0],
				"Holding_id": holdingIds[0],
				"Total":      "One Thousand and Fifty",
			},
			ExpectedResponse:   routeTester.ExpectedUpdateSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
