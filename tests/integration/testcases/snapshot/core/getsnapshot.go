package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotTestCases(t *testing.T, snapshotId, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		{
			Title:       "200",
			ParameterId: snapshotId,
			ExpectedResponse: snapshotTester.ExpectedGetSnapshotResponse{
				Total:         CoreSnapshotTotal,
				WeightedErPct: CoreWeightedEr,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        snapshotId,
			ReplacementToken:   tok401,
			ExpectedResponse:   snapshotTester.ExpectedGetSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        9999,
			ExpectedResponse:   snapshotTester.ExpectedGetSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
