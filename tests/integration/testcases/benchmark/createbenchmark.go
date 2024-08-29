package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetCreateBenchmarkTestCases(t *testing.T, userId int, email string) []shared.PostTestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.PostTestCase{
		// ---- 200 ----
		{
			Title: "200",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:        "Classic Bogleheads Portfolio",
				Description: "The classic 60/40 split",
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
				Std_dev_pct:     3.4,
				Real_return_pct: 6.8,
				Drawdown_yrs:    10,
				Is_deprecated:   true,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "200 Valid Duplicate",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:        "Classic Bogleheads Portfolio",
				Description: "The classic 60/40 split",
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
				Std_dev_pct:     3.4,
				Real_return_pct: 6.8,
				Drawdown_yrs:    10,
				Is_deprecated:   false,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "200 No Description",
			Payload: benchmark.CreateBenchmarkPayload{
				Name: "Golden Butterfly",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  20,
					},
					{
						Category: "DSCV",
						Percent:  20,
					},
					{
						Category: "LTB",
						Percent:  20,
					},
					{
						Category: "GOLD",
						Percent:  20,
					},
					{
						Category: "STB",
						Percent:  20,
					},
				},
				Std_dev_pct:     0.5,
				Real_return_pct: 6.5,
				Drawdown_yrs:    3,
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},
		{
			Title: "200 Name & Payload",
			Payload: benchmark.CreateBenchmarkPayload{
				Name: "Permanent Portfolio",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  25,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "CASH",
						Percent:  25,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusCreated,
		},

		// ---- 401, 409 ----
		{
			Title:              "401",
			Payload:            benchmark.CreateBenchmarkPayload{},
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title: "409",
			Payload: benchmark.CreateBenchmarkPayload{
				Name: "Classic Bogleheads Portfolio",
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
			},
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 ----
		{
			Title: "400 Bad Drawdown",
			Payload: map[string]any{
				"Name": "Benchmark 1",
				"Asset_allocation": []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  25,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "STB",
						Percent:  25,
					},
				},
				"Drawdown_yrs": 10.1,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Real Return",
			Payload: map[string]any{
				"Name": "Benchmark 1",
				"Asset_allocation": []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  25,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "STB",
						Percent:  25,
					},
				},
				"Real_return_pct": "one",
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Name",
			Payload: map[string]any{
				"Name": 1,
				"Asset_allocation": []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  25,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "STB",
						Percent:  25,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Name",
			Payload: benchmark.CreateBenchmarkPayload{
				Name: "",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  25,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "STB",
						Percent:  25,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Sum 1",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:        "Classic Bogleheads Portfolio",
				Description: "The classic 60/40 split",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  50,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				Std_dev_pct:     1,
				Real_return_pct: 7,
				Drawdown_yrs:    10,
				Is_deprecated:   false,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Bad Sum 2",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:        "Golden Butterfly",
				Description: "A juiced up Permanent Portfolio",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  35,
					},
					{
						Category: "LTB",
						Percent:  25,
					},
					{
						Category: "GOLD",
						Percent:  25,
					},
					{
						Category: "STB",
						Percent:  25,
					},
				},
				Std_dev_pct:     0.5,
				Real_return_pct: 6.5,
				Drawdown_yrs:    3,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 No Asset Allocation",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:            "Permanent Portfolio",
				Description:     "Harry Browne's all-weather portfolio",
				Std_dev_pct:     0.5,
				Real_return_pct: 6.5,
				Drawdown_yrs:    3,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Asset Category 1",
			Payload: benchmark.CreateBenchmarkPayload{
				Name:        "Classic Bogleheads Portfolio",
				Description: "The classic 60/40 split",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TEST",
						Percent:  50,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				Std_dev_pct:     1,
				Real_return_pct: 7,
				Drawdown_yrs:    10,
				Is_deprecated:   false,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Asset Category 2",
			Payload: map[string]any{
				"Name": "Permanent Portfolio",
				"Asset_allocation": []map[string]any{
					{
						"Category": 35,
						"Percent":  25,
					},
					{
						"Category": "LTB",
						"Percent":  25,
					},
					{
						"Category": "GOLD",
						"Percent":  25,
					},
					{
						"Category": "STB",
						"Percent":  25,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Asset Category 3",
			Payload: map[string]any{
				"Name": "Permanent Portfolio",
				"Asset_allocation": []map[string]any{
					{
						"Category": "TSM",
						"Percent":  "TSM",
					},
					{
						"Category": "LTB",
						"Percent":  25,
					},
					{
						"Category": "GOLD",
						"Percent":  25,
					},
					{
						"Category": "STB",
						"Percent":  25,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Asset Category 4",
			Payload: benchmark.CreateBenchmarkPayload{
				Name: "Total Stock Market",
				Asset_allocation: []types.AssetAllocationPct{
					{
						Category: "TSM",
						Percent:  100,
					},
					{
						Category: "ITB",
						Percent:  0,
					},
				},
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
