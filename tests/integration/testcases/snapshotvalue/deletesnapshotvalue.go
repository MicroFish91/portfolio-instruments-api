package snapshotvalue

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func DeleteSnapshotValueTestCases(t *testing.T, sid, svid, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:              "401",
			ParameterId:        sid,
			ParameterId2:       svid,
			ReplacementToken:   tok401,
			ExpectedResponse:   routeTester.ExpectedDeleteSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "200",
			ParameterId:        sid,
			ParameterId2:       svid,
			ExpectedResponse:   routeTester.ExpectedDeleteSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "404",
			ParameterId:        sid,
			ParameterId2:       svid,
			ExpectedResponse:   routeTester.ExpectedDeleteSnapshotValueResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
