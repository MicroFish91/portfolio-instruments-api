package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByInstitutionTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotByInstitutionResponse{
			Institutions: []string{
				"Vanguard",
				"Fidelity",
				"Schwab",
			},
			Totals: []float64{
				159676.58,
				141051.39,
				206380.13,
			},
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
