package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func CreateHoldingTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 200 ----
		{
			Title: "201 Ticker",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Ticker Deprecated",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund 2",
				Ticker:            "VTSAX2",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
				Is_deprecated:     true,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 No Ticker",
			Payload: holding.CreateHoldingPayload{
				Name:           "Bank01",
				Ticker:         "",
				Asset_category: "CASH",
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "201 Fixed Income",
			Payload: holding.CreateHoldingPayload{
				Name:              "9128285M8",
				Asset_category:    "LTB",
				Maturation_date:   "07/01/2054",
				Interest_rate_pct: 4.6,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},

		// ---- 401, 409 ----
		{
			Title:            "401",
			ReplacementToken: tok401,
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title: "409 Duplicate Name",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAXX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},
		{
			Title: "409 Duplicate Ticker + Deprecated",
			Payload: holding.CreateHoldingPayload{
				Name:              "Total Stock Market",
				Ticker:            "VTSAX2",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 ----
		{
			Title: "400 No Name",
			Payload: holding.CreateHoldingPayload{
				Name:              "",
				Ticker:            "VTSAX",
				Asset_category:    "TSM",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Name",
			Payload: map[string]any{
				"Name":              5,
				"Ticker":            "VTSAX",
				"Asset_category":    "TSM",
				"Expense_ratio_pct": 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Category",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Category",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "OIL",
				Expense_ratio_pct: 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Category 2",
			Payload: map[string]any{
				"Name":              "Vanguard Total Stock Market Index Fund",
				"Ticker":            "VTSAX",
				"Asset_category":    5,
				"Expense_ratio_pct": 0.04,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Expense Ratio 1",
			Payload: holding.CreateHoldingPayload{
				Name:              "Vanguard Total Stock Market Index Fund",
				Ticker:            "VTSAX",
				Asset_category:    "OIL",
				Expense_ratio_pct: 0.098174,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Expense Ratio 2",
			Payload: map[string]any{
				"Name":              "Vanguard Total Stock Market Index Fund",
				"Ticker":            "VTSAX",
				"Asset_category":    "TSM",
				"Expense_ratio_pct": "eight-six",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Maturation Date",
			Payload: holding.CreateHoldingPayload{
				Name:            "9128285M9",
				Asset_category:  "STB",
				Maturation_date: "07/01/26",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Interest Rate",
			Payload: holding.CreateHoldingPayload{
				Name:              "9128285M9",
				Asset_category:    "STB",
				Maturation_date:   "07/01/2026",
				Interest_rate_pct: 120,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
