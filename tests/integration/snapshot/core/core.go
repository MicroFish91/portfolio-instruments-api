package core

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	coreSnapshotTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshot/core"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
)

var (
	ss_core_testuser types.User
	ss_core_token    string

	ss_core_benchmarkid int
	ss_core_accountids  []int
	ss_core_holdingids  []int

	ss_core_snapid int
	ss_core_svids  []int
)

func CoreSnapshotScenarioTests(t *testing.T) {
	t.Run("Setup", coreSnapshotSetup)
	t.Run("POST://api/v1/snapshots", createSnapshotTest)
}

func coreSnapshotSetup(t *testing.T) {
	ss_core_testuser, ss_core_token = integration.NewUserSetup(t)
	ss_core_benchmarkid = createCoreSnapshotBenchmark(t, ss_core_token, ss_core_testuser.User_id)
	ss_core_accountids = createCoreSnapshotAccounts(t, ss_core_token, ss_core_testuser.User_id)
	ss_core_holdingids = createCoreSnapshotHoldings(t, ss_core_token, ss_core_testuser.User_id)
}

func createSnapshotTest(t *testing.T) {
	testcases := coreSnapshotTestCases.GetCreateSnapshotTestCases(t, ss_core_benchmarkid, ss_core_accountids, ss_core_holdingids, ss_core_testuser.User_id, ss_core_testuser.Email)

	for _, tc := range testcases {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedCreateSnapshotResponse)
			if !ok {
				t.Fatal("invalid ExpectedCreateSnapshotResponse")
			}

			snapid, svids := snapshotTester.TestCreateSnapshot(
				t2,
				tc.Payload,
				tok,
				expected,
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
			)

			if ss_core_snapid == 0 || ss_core_svids == nil {
				ss_core_snapid = snapid
				ss_core_svids = svids
			}
		})
	}
}
