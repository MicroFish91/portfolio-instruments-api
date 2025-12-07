package tests

import (
	"os"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/servicesuites"
	snapshotSuite "github.com/MicroFish91/portfolio-instruments-api/tests/integration/servicesuites/snapshot"
	snapshotValueSuite "github.com/MicroFish91/portfolio-instruments-api/tests/integration/servicesuites/snapshotvalue"
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
	t.Run("V1", TestV1)

	t.Run("V2-Integration", func(t2 *testing.T) {
		t2.Run("Users-Auth", servicesuites.TestUserService)
		t2.Run("Benchmarks", servicesuites.TestBenchmarkService)
		t2.Run("Accounts", servicesuites.TestAccountService)
		t2.Run("Holdings", servicesuites.TestHoldingService)
		t2.Run("Snapshots", snapshotSuite.TestSnapshotService)
		t2.Run("SnapshotValue", snapshotValueSuite.TestSnapshotValueService)
	})
}
