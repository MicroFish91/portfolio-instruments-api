package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByAccountsTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshotTester.ExpectedGetSnapshotByAccountResponse{
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
