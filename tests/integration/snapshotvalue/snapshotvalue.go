package snapshotvalue

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	snapshotValueTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshotvalue"
	snapshotValueTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/snapshotvalue"
)

var (
	svs_token    string
	svs_testuser types.User

	svs_snapshotid  int
	svs_benchmarkid int
	svs_accountids  []int
	svs_holdingids  []int
	svs_svids       []int
)

func TestSnapshotValueService(t *testing.T) {
	t.Run("Setup", snapshotValueServiceSetup)
	t.Run("POST://api/v1/snapshots/:sid/values", createSnapshotValueTests)
	t.Run("GET://api/v1/snapshots/:sid/values", getSnapshotValuesTests)
	t.Run("GET://api/v1/snapshots/:sid/values/:svid", getSnapshotValueTests)
}

func snapshotValueServiceSetup(t *testing.T) {
	svs_testuser, svs_token = integration.NewUserSetup(t)
	svs_benchmarkid = createSvsBenchmark(t, svs_token, svs_testuser.User_id)
	svs_accountids = createSvsAccounts(t, svs_token, svs_testuser.User_id)
	svs_holdingids = createSvsHoldings(t, svs_token, svs_testuser.User_id)
	svs_snapshotid = createSvsSnapshot(t, svs_accountids, svs_holdingids, svs_benchmarkid, svs_token, svs_testuser.User_id)
}

func createSnapshotValueTests(t *testing.T) {
	for _, tc := range snapshotValueTestCases.CreateSnapshotValueTestCases(t, svs_accountids, svs_holdingids, svs_snapshotid, svs_testuser.User_id, svs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := svs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			svid := snapshotValueTester.TestCreateSnapshotValue(
				t2,
				tc.Payload,
				tok,
				tc.ParameterId,
				svs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
			svs_svids = append(svs_svids, svid)
		})
	}
}

func getSnapshotValuesTests(t *testing.T) {
	for _, tc := range snapshotValueTestCases.GetSnapshotValuesTestCases(t, svs_snapshotid, svs_testuser.User_id, svs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := svs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			expectedResponse, ok := tc.ExpectedResponse.(snapshotValueTester.ExpectedGetSnapshotValuesResponse)
			if !ok {
				t.Fatal("invalid ExpectedGetSnapshotValuesResponse")
			}

			snapshotValueTester.TestGetSnapshotValues(
				t2,
				tok,
				tc.ParameterId,
				svs_testuser.User_id,
				tc.ExpectedStatusCode,
				expectedResponse,
			)
		})
	}
}

func getSnapshotValueTests(t *testing.T) {
	for _, tc := range snapshotValueTestCases.GetSnapshotValueTestCases(t, svs_snapshotid, svs_svids[0], svs_testuser.User_id, svs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := svs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			snapshotValueTester.TestGetSnapshotValue(
				t2,
				tc.ParameterId,
				tc.ParameterId2,
				tok,
				svs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}
