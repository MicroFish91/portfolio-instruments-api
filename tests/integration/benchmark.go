package integration

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	benchmarkTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/benchmark"
	benchmarkTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/benchmark"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

var (
	bs_testuser types.User
	bs_token    string

	benchmarkId int
	bs_depidx   int
)

func TestBenchmarkService(t *testing.T) {
	t.Parallel()

	t.Run("Setup", benchmarkServiceSetup)
	t.Run("POST://api/v1/benchmarks", createBenchmarkTests)
	t.Run("GET://api/v1/benchmarks", getBenchmarksTests)
	t.Run("GET://api/v1/benchmarks/:id", getBenchmarkTests)

}

func benchmarkServiceSetup(t *testing.T) {
	bs_testuser, bs_token = newUserSetup(t)
}

func createBenchmarkTests(t *testing.T) {
	for _, tc := range benchmarkTestCases.GetCreateBenchmarkTestCases(t, bs_testuser.User_id, bs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := bs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			id := benchmarkTester.TestCreateBenchmark(
				t2,
				tc.Payload,
				tok,
				bs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
			if benchmarkId == 0 && id != 0 {
				benchmarkId = id
			}
		})
	}
}

func getBenchmarksTests(t *testing.T) {
	// Get benchmarks test setup
	t.Run("Setup", func(t2 *testing.T) {
		for i := 1; i <= 26; i += 1 {
			benchmarkTester.TestCreateBenchmark(
				t2,
				benchmark.CreateBenchmarkPayload{
					Name: fmt.Sprintf("Classic Portfolio %d", i),
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
					Std_dev_pct:     7.8,
					Real_return_pct: 6.4,
					Drawdown_yrs:    5,
					Is_deprecated:   utils.GetRotatingDeprecation(&bs_depidx),
				},
				bs_token,
				bs_testuser.User_id,
				fiber.StatusCreated,
			)
		}
	})

	// Get benchmarks tests
	for _, tc := range benchmarkTestCases.GetBenchmarksTestCases(t, bs_testuser.User_id, bs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			response, ok := tc.ExpectedResponse.(benchmarkTestCases.GetBenchmarksExpectedResponse)
			if !ok {
				t.Fatal("invalid GetBenchmarksExpectedResponse")
			}

			tok := bs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			benchmarkTester.TestGetBenchmarks(
				t2,
				tc.Route,
				tok,
				bs_testuser.User_id,
				tc.ExpectedStatusCode,
				response,
			)
		})
	}
}

func getBenchmarkTests(t *testing.T) {
	for _, tc := range benchmarkTestCases.GetBenchmarkTestCases(t, benchmarkId, bs_testuser.User_id, bs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := bs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			benchmarkTester.TestGetBenchmark(
				t2,
				tc.ParameterId,
				tok,
				bs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}
