package advanced

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByTaxShelterTestCase(*testing.T) testcases.TestCase {
	return testcases.TestCase{
		Title: "200",
		ExpectedResponse: routeTester.ExpectedGetSnapshotByTaxShelterResponse{
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
