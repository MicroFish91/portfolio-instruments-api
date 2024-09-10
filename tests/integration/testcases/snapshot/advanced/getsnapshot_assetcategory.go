package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByAssetCategoryTestCase(*testing.T) shared.GetTestCase {
	return shared.GetTestCase{
		Title: "200",
		ExpectedResponse: snapshot.ExpectedGetSnapshotByAssetCategoryResponse{
			HoldingNames: []string{
				"CASH",
				"OTHER",
				"STB",
				"GOLD",
				"TSM",
				"LTB",
				"ITB",
				"DSCV",
			},
			Totals: []float64{
				76120.43,
				93682.02,
				38459.64,
				56417.36,
				41711.18,
				50848.65,
				60268.38,
				89600.44,
			},
		},
		ExpectedStatusCode: fiber.StatusOK,
	}
}