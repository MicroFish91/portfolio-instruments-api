package core

import (
	"testing"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	coreSnapshotTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshot/core"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshot"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/user"
	"github.com/gofiber/fiber/v3"
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
	t.Run("POST://api/v2/snapshots", createSnapshotTest)
	t.Run("GET://api/v2/snapshots", getSnapshotsTests)
	t.Run("GET://api/v2/snapshots/:id", getSnapshotTest)
	t.Run("GET://api/v2/snapshots/:id/rebalance", getSnapshotRebalanceTest)
	t.Run("PUT://api/v2/snapshots/:id", updateSnapshotTest)
	t.Run("DEL://api/v2/snapshots/:id", deleteSnapshotTest)
	t.Run("Cleanup", snapshotServiceCleaner)
}

func coreSnapshotSetup(t *testing.T) {
	ss_core_testuser, ss_core_token = integration.NewUserSetup(t)
	ss_core_benchmarkid = createCoreSnapshotBenchmark(t, ss_core_token, ss_core_testuser.User_id)
	ss_core_accountids = createCoreSnapshotAccounts(t, ss_core_token, ss_core_testuser.User_id)
	ss_core_holdingids = createCoreSnapshotHoldings(t, ss_core_token, ss_core_testuser.User_id)
}

func createSnapshotTest(t *testing.T) {
	for _, tc := range coreSnapshotTestCases.CreateSnapshotTestCases(t, ss_core_benchmarkid, ss_core_accountids, ss_core_holdingids, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedCreateSnapshotResponse)
			if !ok {
				t2.Fatal("invalid ExpectedCreateSnapshotResponse")
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

func getSnapshotsTests(t *testing.T) {
	t.Run("Setup", func(t2 *testing.T) {
		snapDate := time.Now()
		snapDate = snapDate.AddDate(-1, -3, 0)

		for i := 0; i <= 25; i += 1 {
			snapDate = snapDate.AddDate(0, -1, 0)

			snapshotTester.TestCreateSnapshot(
				t2,
				snapshot.CreateSnapshotPayload{
					Snap_date: snapDate.Format("01/02/2006"),
					Snapshot_values: []snapshotvalue.CreateSnapshotValuePayload{
						{
							Account_id: ss_core_accountids[0],
							Holding_id: ss_core_holdingids[0],
							Total:      1000.00,
						},
					},
					Benchmark_id: ss_core_benchmarkid,
				},
				ss_core_token,
				snapshotTester.ExpectedCreateSnapshotResponse{
					Total:         1000,
					WeightedErPct: 0.3,
				},
				ss_core_testuser.User_id,
				fiber.StatusCreated,
			)
		}
	})

	for _, tc := range coreSnapshotTestCases.GetSnapshotsTestCases(t, ss_core_snapid, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expectedResponse, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotsResponse)
			if !ok {
				t2.Fatal("invalid ExpectedGetSnapshotsResponse")
			}

			snapshotTester.TestGetSnapshots(
				t2,
				tc.Route,
				tok,
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
				expectedResponse,
			)
		})
	}
}

func getSnapshotTest(t *testing.T) {
	for _, tc := range coreSnapshotTestCases.GetSnapshotTestCases(t, ss_core_snapid, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotResponse)
			if !ok {
				t2.Fatal("invalid ExpectedGetSnapshotResponse")
			}

			snapshotTester.TestGetSnapshot(
				t2,
				tc.ParameterId,
				tok,
				snapshotTester.ExpectedGetSnapshotResponse{
					AccountIds:    []int{ss_core_accountids[0], ss_core_accountids[1]},
					HoldingIds:    ss_core_holdingids,
					Total:         expected.Total,
					WeightedErPct: expected.WeightedErPct,
				},
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func getSnapshotRebalanceTest(t *testing.T) {
	for _, tc := range coreSnapshotTestCases.GetSnapshotRebalanceTestCases(t, ss_core_snapid, ss_core_snapid+1, ss_core_snapid+2, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotRebalanceResponse)
			if !ok {
				t2.Fatal("invalid ExpectedGetSnapshotRebalanceResponse")
			}

			snapshotTester.TestGetSnapshotRebalance(
				t2,
				tc.ParameterId,
				tok,
				expected,
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func updateSnapshotTest(t *testing.T) {
	for _, tc := range coreSnapshotTestCases.UpdateSnapshotTestCases(t, ss_core_snapid, ss_core_benchmarkid, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			snapshotTester.TestUpdateSnapshot(
				t2,
				tc.ParameterId,
				tc.Payload,
				tok,
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func deleteSnapshotTest(t *testing.T) {
	for _, tc := range coreSnapshotTestCases.DeleteSnapshotTestCases(t, ss_core_snapid, ss_core_testuser.User_id, ss_core_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := ss_core_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			snapshotTester.TestDeleteSnapshot(
				t2,
				tc.ParameterId,
				tok,
				ss_core_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func snapshotServiceCleaner(t *testing.T) {
	userTester.TestDeleteUser(t, "", ss_core_token, ss_core_testuser.User_id, fiber.StatusOK)
}
