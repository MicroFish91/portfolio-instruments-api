package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/benchmark"
	benchmarkTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/benchmark"
)

var (
	bs_testuser types.User
	bs_token    string
)

func TestBenchmarkService(t *testing.T) {
	t.Parallel()

	t.Run("Setup", benchmarkServiceSetup)
	t.Run("POST://api/v1/benchmarks", createBenchmarkTests)
}

func benchmarkServiceSetup(t *testing.T) {
	bs_testuser, bs_token = newUserSetup(t)
}

func createBenchmarkTests(t *testing.T) {
	for _, tc := range benchmark.GetCreateBenchmarkTestCases(t, bs_testuser.User_id, bs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := bs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			benchmarkTester.TestCreateBenchmark(
				t2,
				tc.Payload,
				tok,
				bs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}
