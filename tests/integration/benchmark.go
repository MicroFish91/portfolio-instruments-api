package integration

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
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
	fmt.Println(bs_testuser, bs_token)
}
