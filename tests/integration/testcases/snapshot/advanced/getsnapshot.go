package advanced

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotTestCase(*testing.T) testcases.TestCase {
	return testcases.TestCase{
		Title: "200",
		ExpectedResponse: routeTester.ExpectedGetSnapshotResponse{
			Total:         AdvancedSnapshotTotal,
			WeightedErPct: AdvancedSnapshotExpenseRatio,
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
