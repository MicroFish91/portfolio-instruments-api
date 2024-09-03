package advanced

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	accountTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/account"
	benchmarkTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/benchmark"
	holdingTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/holding"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

// Create benchmark
func createAdvancedSnapshotBenchmark(t *testing.T, token string, userId int) int {
	var id int

	t.Run("Benchmark", func(t2 *testing.T) {
		id = benchmarkTester.TestCreateBenchmark(
			t2,
			benchmark.CreateBenchmarkPayload{
				Name: "Golden Butterfly Portfolio",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  20,
					},
					{
						Category: "LTB",
						Percent:  20,
					},
					{
						Category: "CASH",
						Percent:  20,
					},
					{
						Category: "GOLD",
						Percent:  20,
					},
					{
						Category: "DSCV",
						Percent:  20,
					},
				},
				Is_deprecated: false,
			},
			token,
			userId,
			fiber.StatusCreated,
		)
	})

	return id
}

// Create accounts
func createAdvancedSnapshotAccounts(t *testing.T, token string, userId int) []int {
	var i = 1
	var ids []int

	for _, institution := range []string{"Vanguard", "Fidelity", "Schwab"} {
		for _, ts := range []string{"TAXABLE", "ROTH", "TRADITIONAL"} {
			name := fmt.Sprintf("Account%d", i)

			t.Run(name, func(t2 *testing.T) {
				id := accountTester.TestCreateAccount(
					t2,
					account.CreateAccountPayload{
						Name:          name,
						Tax_shelter:   ts,
						Institution:   institution,
						Is_deprecated: false,
					},
					token,
					userId,
					fiber.StatusCreated,
				)

				ids = append(ids, id)
				i += 1
			})
		}
	}

	return ids
}

// Create holdings
var assets = []struct {
	ticker          string
	asset_category  string
	expense_ratio   float32
	maturation_date string
	interest_rate   float32
}{
	{
		ticker:         "AAAA",
		asset_category: "TSM",
		expense_ratio:  0.3,
	},
	{
		asset_category:  "LTB",
		maturation_date: utils.Calc_target_date(28, 6),
		interest_rate:   5.8,
	},
	{
		asset_category: "CASH",
	},
	{
		ticker:         "BBBB",
		asset_category: "GOLD",
		expense_ratio:  0.82,
	},
	{
		ticker:         "CCCC",
		asset_category: "DSCV",
		expense_ratio:  0.64,
	},
	{
		ticker:         "DDDD",
		asset_category: "OTHER",
		expense_ratio:  0.58,
	},
	{
		asset_category:  "STB",
		maturation_date: utils.Calc_target_date(2, 6),
		interest_rate:   3.6,
	},
	{
		ticker:         "EEEE",
		asset_category: "TSM",
		expense_ratio:  0.18,
	},
	{
		asset_category:  "LTB",
		maturation_date: utils.Calc_target_date(26, 0),
		interest_rate:   5.5,
	},
	{
		asset_category: "CASH",
	},
	{
		ticker:         "FFFF",
		asset_category: "GOLD",
		expense_ratio:  0.77,
	},
	{
		ticker:         "GGGG",
		asset_category: "DSCV",
		expense_ratio:  0.9,
	},
	{
		ticker:         "HHHH",
		asset_category: "OTHER",
		expense_ratio:  1.18,
	},
	{
		asset_category:  "ITB",
		maturation_date: utils.Calc_target_date(10, 0),
		interest_rate:   4.4,
	},
}

func createAdvancedSnapshotHoldings(t *testing.T, token string, userId int) []int {
	var i = 1
	var ids []int

	for _, asset := range assets {
		name := fmt.Sprintf("Holding%d", i)

		t.Run(name, func(t2 *testing.T) {
			id := holdingTester.TestCreateHolding(
				t2,
				holding.CreateHoldingPayload{
					Name:              name,
					Ticker:            asset.ticker,
					Asset_category:    asset.asset_category,
					Expense_ratio_pct: asset.expense_ratio,
					Maturation_date:   asset.maturation_date,
					Interest_rate_pct: asset.interest_rate,
					Is_deprecated:     false,
				},
				token,
				userId,
				fiber.StatusCreated,
			)

			ids = append(ids, id)
			i += 1
		})
	}

	return ids
}
