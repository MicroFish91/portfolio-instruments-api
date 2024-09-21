package snapshotvalue

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotValuesTestCases(t *testing.T, snapshotId, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		{
			Title:       "200",
			ParameterId: snapshotId,
			ExpectedResponse: snapshotvalue.ExpectedGetSnapshotValuesResponse{
				SnapshotValues: 3,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        snapshotId,
			ReplacementToken:   tok401,
			ExpectedResponse:   snapshotvalue.ExpectedGetSnapshotValuesResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        99999,
			ExpectedResponse:   snapshotvalue.ExpectedGetSnapshotValuesResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
