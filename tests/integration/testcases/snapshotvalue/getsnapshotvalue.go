package snapshotvalue

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotValueTestCases(t *testing.T, sid, svid, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		{
			Title:              "200",
			ParameterId:        sid,
			ParameterId2:       svid,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:              "401",
			ParameterId:        sid,
			ParameterId2:       svid,
			ReplacementToken:   tok401,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotValuesResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:              "404 sid",
			ParameterId:        99999,
			ParameterId2:       svid,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotValuesResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
		{
			Title:              "404 svid",
			ParameterId:        sid,
			ParameterId2:       99999,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotValuesResponse{},
			ExpectedStatusCode: fiber.StatusNotFound,
		},
	}
}
