package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetUpdateHoldingTestCases(t *testing.T, holdingId int, userId int, email string) []shared.PutTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.PutTestCase{
		// ---- 200 ----
		{
			Title:       "200 Ticker",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Fidelity Total Market Index Fund",
				Ticker:            "FSKAX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 No Ticker",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:           "Bank02",
				Ticker:         "",
				Asset_category: "CASH",
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Fixed Income",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "9138285M8",
				Asset_category:    "LTB",
				Maturation_date:   "07/01/2054",
				Interest_rate_pct: 4.6,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Is Deprecated",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Fidelity Small Cap Value Index Fund",
				Ticker:            "FSSNAX",
				Asset_category:    "DSCV",
				Expense_ratio_pct: 0.04,
				Is_deprecated:     true,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},

		// ---- 400, 404 ----
		{
			Title:            "401",
			ParameterId:      holdingId,
			ReplacementToken: tok401,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Fidelity Small Cap Value Index Fund",
				Ticker:            "FSSNAX",
				Asset_category:    "DSCV",
				Expense_ratio_pct: 0.04,
				Is_deprecated:     true,
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "404",
			ParameterId: 9999,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Fidelity Small Cap Value Index Fund",
				Ticker:            "FSSNAX",
				Asset_category:    "DSCV",
				Expense_ratio_pct: 0.04,
				Is_deprecated:     true,
			},
			ExpectedStatusCode: fiber.StatusNotFound,
		},

		// ---- 400 ----
		{
			Title:       "400 No Name",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "",
				Ticker:            "VTSAX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Name",
			ParameterId: holdingId,
			Payload: map[string]any{
				"Name":              5,
				"Ticker":            "VTSAX",
				"Asset_category":    "TSM",
				"Expense_ratio_pct": 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 No Category",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Category",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "OIL",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Category 2",
			ParameterId: holdingId,
			Payload: map[string]any{
				"Name":              "Vanguard Total Stock Market Index Fund",
				"Ticker":            "VTSAX",
				"Asset_category":    5,
				"Expense_ratio_pct": 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Expense Ratio 1",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "OIL",
				Expense_ratio_pct: 0.098174,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Expense Ratio 2",
			ParameterId: holdingId,
			Payload: map[string]any{
				"Name":              "Vanguard Total Stock Market Index Fund",
				"Ticker":            "VTSAX",
				"Asset_category":    "TSM",
				"Expense_ratio_pct": "eight-six",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Maturation Date",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:            "9128285M9",
				Asset_category:  "STB",
				Maturation_date: "07/01/26",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:       "400 Bad Interest Rate",
			ParameterId: holdingId,
			Payload: holding.UpdateHoldingPayload{
				Name:              "9128285M9",
				Asset_category:    "STB",
				Maturation_date:   "07/01/2026",
				Interest_rate_pct: 120,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
