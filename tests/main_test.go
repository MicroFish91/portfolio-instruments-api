package tests

import (
	"os"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	snapIntegration "github.com/MicroFish91/portfolio-instruments-api/tests/integration/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testserver"
)

func TestMain(m *testing.M) {
	tsw := testserver.GetTestServerWrapper()
	defer tsw.Shutdown()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestApi(t *testing.T) {
	t.Run("Ping", TestPing)

	t.Run("Integration", func(t2 *testing.T) {
		t2.Run("Users-Auth", integration.TestUserService)
		t2.Run("Benchmarks", integration.TestBenchmarkService)
		t2.Run("Accounts", integration.TestAccountService)
		t2.Run("Holdings", integration.TestHoldingService)
		t2.Run("Snapshots", snapIntegration.TestSnapshotService)
		t2.Run("SnapshotValue", integration.TestSnapshotValueService)
	})
}
