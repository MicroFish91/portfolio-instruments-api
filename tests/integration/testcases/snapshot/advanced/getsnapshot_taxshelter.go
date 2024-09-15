package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByTaxShelterTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotByTaxShelterResponse{
			TaxShelters: []string{
				"TAXABLE",
				"ROTH",
				"TRADITIONAL",
			},
			Totals: []float64{
				144891.62,
				189163.79,
				173052.69,
			},
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
