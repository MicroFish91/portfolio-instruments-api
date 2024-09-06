package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshot.ExpectedGetSnapshotResponse{
			Total:         AdvancedSnapshotTotal,
			WeightedErPct: AdvancedSnapshotExpenseRatio,
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
