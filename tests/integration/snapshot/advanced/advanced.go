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
	t.Run("POST://api/v1/snapshots", createSnapshotTest)
	t.Run("GET://api/v1/snapshots/:id", getSnapshotTest)
	t.Run("GET://api/v1/snapshots/:id?group_by=ACCOUNT_NAME", getSnapshotByAccountNameTest)
	t.Run("GET://api/v1/snapshots/:id?group_by=ACCOUNT_INSTITUTION", getSnapshotByInstitutionTest)
	t.Run("GET://api/v1/snapshots/:id?group_by=ASSET_CATEGORY", getSnapshotByAssetCategoryTest)
}

func advancedSnapshotSetup(t *testing.T) {
	ss_adv_testuser, ss_adv_token = integration.NewUserSetup(t)
	ss_adv_benchmarkid = createAdvancedSnapshotBenchmark(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_accountids = createAdvancedSnapshotAccounts(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_holdingids = createAdvancedSnapshotHoldings(t, ss_adv_token, ss_adv_testuser.User_id)
}

func createSnapshotTest(t *testing.T) {
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

func getSnapshotTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotResponse")
	}
	expected.AccountIds = ss_adv_accountids
	expected.HoldingIds = ss_adv_holdingids

	snapshotTester.TestGetSnapshot(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}

func getSnapshotByAccountNameTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotByAccountsTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotByAccountResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotByAccountResponse")
	}

	snapshotTester.TestGetSnapshotByAccount(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}

func getSnapshotByAssetCategoryTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotByAssetCategoryTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotByAssetCategoryResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotByAssetCategoryResponse")
	}

	snapshotTester.TestGetSnapshotByAssetCategory(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}

func getSnapshotByInstitutionTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotByInstitutionTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotByInstitutionResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotByInstitutionResponse")
	}

	snapshotTester.TestGetSnapshotByInstitution(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}
