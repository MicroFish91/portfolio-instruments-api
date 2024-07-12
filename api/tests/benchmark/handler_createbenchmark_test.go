package benchmark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/tests/mocks"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type CreateBenchmarkTestCases struct {
	Title       string
	Payloads    []map[string]any
	StatusCode  int
	ExcludeAuth bool
}

var createBenchmarkTestCases = []CreateBenchmarkTestCases{
	// 202 - Created
	{
		Title: "should successfully create resource when given valid payload",
		Payloads: []map[string]any{
			{
				"Name":        "Classic Bogleheads Portfolio",
				"Description": "The classic 60/40 split",
				"Asset_allocation": []types.AssetAllocation{
					{
						Category: "TSM",
						Percent:  60,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				"Std_dev_pct":     1,
				"Real_return_pct": 7,
				"Drawdown_yrs":    10,
				"Is_deprecated":   true,
			},
			{
				"Name":        "Golden Butterfly",
				"Description": "A juiced up Permanent Portfolio",
				"Asset_allocation": []types.AssetAllocation{
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
				"Std_dev_pct":     0.5,
				"Real_return_pct": 6.5,
				"Drawdown_yrs":    3,
			},
			{
				"Name": "Permanent Portfolio",
				"Asset_allocation": []types.AssetAllocation{
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
		},
		StatusCode:  fiber.StatusCreated,
		ExcludeAuth: false,
	},

	// 400 - Bad Request
	{
		Title: "should respond with bad request when given invalid payload",
		Payloads: []map[string]any{
			{
				"Name":        "Classic Bogleheads Portfolio",
				"Description": "The classic 60/40 split",
				"Asset_allocation": []types.AssetAllocation{
					{
						Category: "TSM",
						Percent:  50,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				"Std_dev_pct":     1,
				"Real_return_pct": 7,
				"Drawdown_yrs":    10,
				"Is_deprecated":   false,
			},
			{
				"Name":        "Classic Bogleheads Portfolio",
				"Description": "The classic 60/40 split",
				"Asset_allocation": []types.AssetAllocation{
					{
						Category: "TEST",
						Percent:  50,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				"Std_dev_pct":     1,
				"Real_return_pct": 7,
				"Drawdown_yrs":    10,
				"Is_deprecated":   false,
			},
			{
				"Name":        "Golden Butterfly",
				"Description": "A juiced up Permanent Portfolio",
				"Asset_allocation": []types.AssetAllocation{
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
				"Std_dev_pct":     0.5,
				"Real_return_pct": 6.5,
				"Drawdown_yrs":    3,
			},
			{
				"Name": 1,
				"Asset_allocation": []types.AssetAllocation{
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
			},
			{
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
			{
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
			{
				"Asset_allocation": []types.AssetAllocation{
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
			},
			{
				"Name":            "Permanent Portfolio",
				"Description":     "Harry Browne's all-weather portfolio",
				"Std_dev_pct":     0.5,
				"Real_return_pct": 6.5,
				"Drawdown_yrs":    3,
			},
			{
				"Name": "Total Stock Market",
				"Asset_allocation": []types.AssetAllocation{
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
		},
		StatusCode:  400,
		ExcludeAuth: false,
	},

	// 401 - Unauthorized
	{
		Title: "should respond with unauthorized when missing auth token",
		Payloads: []map[string]any{
			{
				"Name":        "Classic Bogleheads Portfolio",
				"Description": "The classic 60/40 split",
				"Asset_allocation": []types.AssetAllocation{
					{
						Category: "TSM",
						Percent:  60,
					},
					{
						Category: "ITB",
						Percent:  40,
					},
				},
				"Std_dev_pct":     1,
				"Real_return_pct": 7,
				"Drawdown_yrs":    10,
				"Is_deprecated":   true,
			},
		},
		StatusCode:  fiber.StatusUnauthorized,
		ExcludeAuth: true,
	},
}

func TestCreateBenchmarkHandler(t *testing.T) {
	app := registerAppWithBenchmarks()

	for _, testCase := range createBenchmarkTestCases {
		t.Run(testCase.Title, func(t *testing.T) {
			for _, payload := range testCase.Payloads {
				body, err := json.Marshal(payload)
				if err != nil {
					t.Fatal(err)
				}

				req, err := http.NewRequest(http.MethodPost, "/api/v1/benchmarks", bytes.NewBuffer(body))
				if err != nil {
					t.Fatal(err)
				}

				req.Header.Set("Content-Type", "application/json")

				if !testCase.ExcludeAuth {
					mocks.MockAuthRequestHeaders(req)
				}

				res, err := app.Test(req)
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(t, testCase.StatusCode, res.StatusCode)
			}
		})
	}
}
