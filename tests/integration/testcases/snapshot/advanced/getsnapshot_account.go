package advanced

import (
	"testing"

	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByAccountsTestCase(*testing.T) testcases.TestCase {
	return testcases.TestCase{
		Title: "200",
		ExpectedResponse: routeTester.ExpectedGetSnapshotByAccountResponse{
			AccountNames: []string{
				"Account1",
				"Account2",
				"Account3",
				"Account4",
				"Account5",
				"Account6",
				"Account7",
				"Account8",
				"Account9",
			},
			Totals: []float64{
				36527.58,
				63578.53,
				59570.47,
				37341.87,
				59384.99,
				44324.53,
				71022.17,
				66200.27,
				69157.69,
			},
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}
