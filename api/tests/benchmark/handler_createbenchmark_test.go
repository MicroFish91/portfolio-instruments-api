package benchmark

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/tests/mocks"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

type CreateBenchmarkTestCases struct {
	Title       string
	Payloads    []benchmark.CreateBenchmarkPayload
	StatusCode  int
	ExcludeAuth bool
}

var createBenchmarkTestCases = []CreateBenchmarkTestCases{
	{
		Title: "should successfully handle benchmark creation",
		Payloads: []benchmark.CreateBenchmarkPayload{
			{
				Name:        "Golden Butterfly",
				Description: "A juiced up Permanent Portfolio",
				Asset_allocation: []types.AssetAllocation{
					{
						Category: "TSM",
						Percent:  50,
					},
					{
						Category: "ITB",
						Percent:  50,
					},
				},
				Std_dev_pct:     0.5,
				Real_return_pct: 0.1,
				Drawdown_yrs:    3,
				Is_deprecated:   false,
			},
		},
		StatusCode: fiber.StatusCreated,
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
