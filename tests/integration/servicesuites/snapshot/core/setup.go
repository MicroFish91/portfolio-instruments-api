package core

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/account"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	accountTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/account"
	benchmarkTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/benchmark"
	holdingTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/holding"
	"github.com/gofiber/fiber/v3"
)

// Create benchmark
func createCoreSnapshotBenchmark(t *testing.T, token string, userId int) int {
	var id int

	t.Run("Benchmark", func(t2 *testing.T) {
		id = benchmarkTester.TestCreateBenchmark(
			t2,
			benchmark.CreateBenchmarkPayload{
				Name: "Bogleheads Portfolio",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  60,
					},
					{
						Category: "ITB",
						Percent:  40,
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
func createCoreSnapshotAccounts(t *testing.T, token string, userId int) []int {
	var i = 1
	var ids []int

	for _, ts := range []string{"TAXABLE", "ROTH", "TRADITIONAL"} {
		name := fmt.Sprintf("Account%d", i)

		t.Run(name, func(t2 *testing.T) {
			id := accountTester.TestCreateAccount(
				t2,
				account.CreateAccountPayload{
					Name:          name,
					Tax_shelter:   ts,
					Institution:   "Fidelity",
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
		ticker:         "BBBB",
		asset_category: "ITB",
		expense_ratio:  0.1,
	},
}

func createCoreSnapshotHoldings(t *testing.T, token string, userId int) []int {
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
