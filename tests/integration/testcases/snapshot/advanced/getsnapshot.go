package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotTestCase(*testing.T) shared.TestCase {
	return shared.TestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotResponse{
			Total:         AdvancedSnapshotTotal,
			WeightedErPct: AdvancedSnapshotExpenseRatio,
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
