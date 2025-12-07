package advanced

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByInstitutionTestCase(*testing.T) testcases.TestCase {
	return testcases.TestCase{
		Title: "200",
		ExpectedResponse: routeTester.ExpectedGetSnapshotByInstitutionResponse{
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
