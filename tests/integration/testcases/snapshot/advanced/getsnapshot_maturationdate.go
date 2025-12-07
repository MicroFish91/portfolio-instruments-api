package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetAdvancedSnapshotByMaturationDateTestCases(*testing.T) []testcases.TestCase {
	return []testcases.TestCase{
		{
			Title: "200 Date 1",
			ExpectedResponse: routeTester.ExpectedGetSnapshotByMaturationDateResponse{
				Resources: []types.MaturationDateResource{
					{
						Account_name:   "Account2",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          2271.85,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account6",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          3483.74,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account7",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          15130.08,
						Skip_rebalance: false,
					},
				},
				Maturation_start: utils.Calc_target_date(27, 6),
				Maturation_end:   "",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Date 2",
			ExpectedResponse: routeTester.ExpectedGetSnapshotByMaturationDateResponse{
				Resources: []types.MaturationDateResource{
					{
						Account_name:   "Account1",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          1092.52,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account2",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          2271.85,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account2",
						Holding_name:   "Holding14",
						Asset_category: "ITB",
						Total:          12096.14,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account2",
						Holding_name:   "Holding9",
						Asset_category: "LTB",
						Total:          13431.37,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account3",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          5880.8,
						Skip_rebalance: true,
					},
					{
						Account_name:   "Account5",
						Holding_name:   "Holding9",
						Asset_category: "LTB",
						Total:          15407.15,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account5",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          6889.24,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account6",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          2240.44,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account6",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          3483.74,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account6",
						Holding_name:   "Holding14",
						Asset_category: "ITB",
						Total:          9448.52,
						Skip_rebalance: true,
					},
					{
						Account_name:   "Account7",
						Holding_name:   "Holding2",
						Asset_category: "LTB",
						Total:          15130.08,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account7",
						Holding_name:   "Holding14",
						Asset_category: "ITB",
						Total:          11453.51,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account8",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          15737.92,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account8",
						Holding_name:   "Holding14",
						Asset_category: "ITB",
						Total:          15411.1,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account9",
						Holding_name:   "Holding7",
						Asset_category: "STB",
						Total:          6618.72,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account9",
						Holding_name:   "Holding9",
						Asset_category: "LTB",
						Total:          1124.46,
						Skip_rebalance: false,
					},
					{
						Account_name:   "Account9",
						Holding_name:   "Holding14",
						Asset_category: "ITB",
						Total:          11859.11,
						Skip_rebalance: true,
					},
				},
				Maturation_start: "",
				Maturation_end:   "",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Date 3",
			ExpectedResponse: routeTester.ExpectedGetSnapshotByMaturationDateResponse{
				Resources: []types.MaturationDateResource{
					{
						Account_name:    "Account1",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           1092.52,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account2",
						Holding_name:    "Holding14",
						Asset_category:  "ITB",
						Maturation_date: "07/01/2010",
						Total:           12096.14,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account3",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           5880.8,
						Skip_rebalance:  true,
					},
					{
						Account_name:    "Account5",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           6889.24,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account6",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           2240.44,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account6",
						Holding_name:    "Holding14",
						Asset_category:  "ITB",
						Maturation_date: "07/01/2010",
						Total:           9448.52,
						Skip_rebalance:  true,
					},
					{
						Account_name:    "Account7",
						Holding_name:    "Holding14",
						Asset_category:  "ITB",
						Maturation_date: "07/01/2010",
						Total:           11453.51,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account8",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           15737.92,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account8",
						Holding_name:    "Holding14",
						Asset_category:  "ITB",
						Maturation_date: "07/01/2010",
						Total:           15411.1,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account9",
						Holding_name:    "Holding7",
						Asset_category:  "STB",
						Maturation_date: "01/01/2003",
						Total:           6618.72,
						Skip_rebalance:  false,
					},
					{
						Account_name:    "Account9",
						Holding_name:    "Holding14",
						Asset_category:  "ITB",
						Maturation_date: "07/01/2010",
						Total:           11859.11,
						Skip_rebalance:  true,
					},
				},
				Maturation_start: "",
				Maturation_end:   utils.Calc_target_date(10, 6),
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
	}
}
