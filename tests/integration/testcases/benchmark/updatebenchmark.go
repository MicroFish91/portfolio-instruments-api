package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func UpdateBenchmarkTestCases(t *testing.T, benchmarkId int, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 200 ----
		{
			Title:       "200",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name:        "Classic Bogleheads Portfolio 3",
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
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 No Description",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name: "Golden Butterfly 2",
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
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Name & Payload",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name: "Permanent Portfolio 2",
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
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Rebalance Threshold 1",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name: "Permanent Portfolio 2",
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
				Rec_rebalance_threshold_pct: 25,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title:       "200 Rebalance Threshold 2",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name: "Permanent Portfolio 2",
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
				Rec_rebalance_threshold_pct: 0,
			},
			ExpectedStatusCode: fiber.StatusOK,
		},

		// ---- 401, 409 ----
		{
			Title:              "401",
			ParameterId:        benchmarkId,
			Payload:            benchmark.UpdateBenchmarkPayload{},
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title:       "409",
			ParameterId: benchmarkId,
			Payload: benchmark.UpdateBenchmarkPayload{
				Name: "claSsiC boglEheadS PortfOlio 2",
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
			Payload: benchmark.UpdateBenchmarkPayload{
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
		{
			Title: "400 Rebalance Threshold",
			Payload: benchmark.UpdateBenchmarkPayload{
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
				Rec_rebalance_threshold_pct: 101,
			},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
