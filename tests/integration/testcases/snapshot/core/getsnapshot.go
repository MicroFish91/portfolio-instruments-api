package core

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotTestCases(t *testing.T, snapshotId, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:       "200",
			ParameterId: snapshotId,
			ExpectedResponse: routeTester.ExpectedGetSnapshotResponse{
				Total:         CoreSnapshotTotal,
				WeightedErPct: CoreWeightedEr,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        snapshotId,
			ReplacementToken:   tok401,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404",
			ParameterId:        9999,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
