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

func TestCreateBenchmarkHandler(t *testing.T) {
	app := registerAppWithBenchmarks()

	t.Run("should successfully handle benchmark creation", func(t *testing.T) {
		payload := benchmark.CreateBenchmarkPayload{
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
		}

		body, err := json.Marshal(payload)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(http.MethodPost, "/api/v1/benchmarks", bytes.NewBuffer(body))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		mocks.MockAuthRequestHeaders(req)

		res, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, fiber.StatusCreated, res.StatusCode)
	})
}
