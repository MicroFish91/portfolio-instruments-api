package advanced

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	advancedSnapshotTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshot/advanced"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
)

var (
	ss_adv_testuser types.User
	ss_adv_token    string

	ss_adv_benchmarkid int
	ss_adv_accountids  []int
	ss_adv_holdingids  []int
	ss_adv_snapid      int
)

func AdvancedSnapshotScenarioTests(t *testing.T) {
	t.Run("Setup", advancedSnapshotSetup)
	t.Run("POST://api/v1/snapshots", createAdvancedSnapshotTests)
}

func advancedSnapshotSetup(t *testing.T) {
	ss_adv_testuser, ss_adv_token = integration.NewUserSetup(t)
	ss_adv_benchmarkid = createAdvancedSnapshotBenchmark(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_accountids = createAdvancedSnapshotAccounts(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_holdingids = createAdvancedSnapshotHoldings(t, ss_adv_token, ss_adv_testuser.User_id)
}

func createAdvancedSnapshotTests(t *testing.T) {
	tc := advancedSnapshotTestCases.GetCreateSnapshotAdvancedTestCase(t, ss_adv_benchmarkid, ss_adv_accountids, ss_adv_holdingids)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedCreateSnapshotResponse)
	if !ok {
		t.Fatal("invalid ExpectedCreateSnapshotResponse")
	}

	ss_adv_snapid, _ = snapshotTester.TestCreateSnapshot(
		t,
		tc.Payload,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}
