package utils

import (
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
)

// GetRotatingEmail
var e_idx int

func GetRotatingEmail() string {
	e_idx += 1
	return fmt.Sprintf("test_user%d@gmail.com", e_idx)
}

// GetRotatingTaxShelter
var taxShelters = []string{
	"TAXABLE",
	"ROTH",
	"TRADITIONAL",
	"HSA",
	"529",
}

func GetRotatingTaxShelter(ts_idx *int) string {
	*ts_idx += 1
	return taxShelters[*ts_idx%len(taxShelters)]
}

// GetRotatingInstitution
var institutions = []string{
	"Vanguard",
	"Fidelity",
	"Schwab",
	"Ameritrade",
}

func GetRotatingInstitution(inst_idx *int) string {
	*inst_idx += 1
	return institutions[*inst_idx%len(institutions)]
}

// GetRotatingDeprecation
func GetRotatingDeprecation(dep_idx *int) bool {
	*dep_idx += 1
	return *dep_idx%10 == 0
}

// GetRotatingExpenseRatio
var expenseRatio = []float32{0.2, 0.4, 0.6, 0.8, 1}

func GetRotatingExpenseRatio(er_idx *int) float32 {
	*er_idx += 1
	return expenseRatio[*er_idx%len(expenseRatio)]
}

// GetRotatingFixedIncome
type MockAsset struct {
	Ticker            string
	Asset_category    types.AssetCategory
	Maturation_date   string
	Interest_rate_pct float32
}

var mockFIAssets = []MockAsset{
	{
		Asset_category: "CASH",
	},
	{
		Asset_category:    "BILLS",
		Maturation_date:   Calc_target_date(0, 6),
		Interest_rate_pct: 2.2,
	},
	{
		Asset_category:    "STB",
		Maturation_date:   Calc_target_date(3, 0),
		Interest_rate_pct: 3.4,
	},
	{
		Asset_category:    "ITB",
		Maturation_date:   Calc_target_date(15, 0),
		Interest_rate_pct: 4.1,
	},
	{
		Asset_category:    "LTB",
		Maturation_date:   Calc_target_date(25, 0),
		Interest_rate_pct: 5.5,
	},
}

func GetRotatingFixedIncome(inc_idx *int) MockAsset {
	*inc_idx += 1
	return mockFIAssets[*inc_idx%len(mockFIAssets)]
}

// GetRotatingMutualFund
var assetCategory = []types.AssetCategory{
	"STB",
	"GOLD",
	"TSM",
	"DSCV",
	"LTB",
}

func GetRotatingMutualFund(mf_idx *int) MockAsset {
	*mf_idx += 1
	return MockAsset{
		Ticker:         fmt.Sprintf("T%d", *mf_idx),
		Asset_category: assetCategory[*mf_idx%len(assetCategory)],
	}
}
