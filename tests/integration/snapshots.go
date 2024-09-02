package integration

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	snapshotTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshot"
)

var (
	ss_testuser types.User
	ss_token    string

	ss_benchmarkid int
	ss_accountids  []int
	ss_holdingids  []int
)

func TestSnapshotService(t *testing.T) {
	t.Run("Setup", snapshotServiceSetup)
	t.Run("POST://api/v1/snapshots", createSnapshotTests)
}

func snapshotServiceSetup(t *testing.T) {
	ss_testuser, ss_token = newUserSetup(t)
	ss_benchmarkid = snapshotTestCases.CreateAdvancedSnapshotBenchmark(t, ss_token, ss_testuser.User_id)
	ss_accountids = snapshotTestCases.CreateAdvancedSnapshotAccounts(t, ss_token, ss_testuser.User_id)
	ss_holdingids = snapshotTestCases.CreateAdvancedSnapshotHoldings(t, ss_token, ss_testuser.User_id)
}

func createSnapshotTests(t *testing.T) {
	//
}
