package advanced

import (
	"math"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/snapshotvalue"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration"
	snapshotTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	snapshotValueTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshotvalue"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/user"
	advancedSnapshotTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/snapshot/advanced"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
)

var (
	ss_adv_testuser types.User
	ss_adv_token    string

	ss_adv_benchmarkid int
	ss_adv_accountids  []int
	ss_adv_holdingids  []int
	ss_adv_svids       []int
	ss_adv_snapid      int

	ss_adv_snapshot types.Snapshot
)

func AdvancedSnapshotScenarioTests(t *testing.T) {
	t.Run("Setup", advancedSnapshotSetup)
	t.Run("POST://api/v2/snapshots", createSnapshotTest)
	t.Run("GET://api/v2/snapshots/:id", getSnapshotTest)
	t.Run("GET://api/v2/snapshots/:id?group_by=ACCOUNT_NAME", getSnapshotByAccountNameTest)
	t.Run("GET://api/v2/snapshots/:id?group_by=ACCOUNT_INSTITUTION", getSnapshotByInstitutionTest)
	t.Run("GET://api/v2/snapshots/:id?group_by=TAX_SHELTER", getSnapshotByTaxShelterTest)
	t.Run("GET://api/v2/snapshots/:id?group_by=ASSET_CATEGORY", getSnapshotByAssetCategoryTest)
	t.Run("GET://api/v2/snapshots/:id?group_by=MATURATION_DATE", getSnapshotByMaturationDateTest)
	t.Run("GET://api/v2/snapshots/:id/rebalance", getSnapshotRebalanceTest)
	t.Run("PUT://api/v2/snapshots/:id/order", updateSnapshotValueOrderTest)

	// snapshot_value
	t.Run("POST://api/v2/snapshots/:id/values/:id", createSnapshotValueTest)
	t.Run("PUT://api/v2/snapshots/:id/values/:id", updateSnapshotValueTest)
	t.Run("DEL://api/v2/snapshots/:id/values/:id", deleteSnapshotValueTest)

	t.Run("DEL://api/v2/snapshots/:id", deleteSnapshotTest)
	t.Run("Cleanup", snapshotServiceCleaner)
}

func advancedSnapshotSetup(t *testing.T) {
	ss_adv_testuser, ss_adv_token = integration.NewUserSetup(t)
	ss_adv_benchmarkid = createAdvancedSnapshotBenchmark(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_accountids = createAdvancedSnapshotAccounts(t, ss_adv_token, ss_adv_testuser.User_id)
	ss_adv_holdingids = createAdvancedSnapshotHoldings(t, ss_adv_token, ss_adv_testuser.User_id)
}

func createSnapshotTest(t *testing.T) {
	tc := advancedSnapshotTestCases.CreateAdvancedSnapshotTestCase(t, ss_adv_benchmarkid, ss_adv_accountids, ss_adv_holdingids)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedCreateSnapshotResponse)
	if !ok {
		t.Fatal("invalid ExpectedCreateSnapshotResponse")
	}

	ss_adv_snapid, ss_adv_svids = snapshotTester.TestCreateSnapshot(
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

	ss_adv_snapshot = snapshotTester.TestGetSnapshot(
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

func getSnapshotByTaxShelterTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotByTaxShelterTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotByTaxShelterResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotByTaxShelterResponse")
	}

	snapshotTester.TestGetSnapshotByTaxShelter(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}

func getSnapshotByMaturationDateTest(t *testing.T) {
	for _, tc := range advancedSnapshotTestCases.GetAdvancedSnapshotByMaturationDateTestCases(t) {
		t.Run(tc.Title, func(t2 *testing.T) {
			expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotByMaturationDateResponse)
			if !ok {
				t2.Fatal("invalid ExpectedGetSnapshotByMaturationDateResponse")
			}

			snapshotTester.TestGetSnapshotByMaturationDate(
				t2,
				ss_adv_snapid,
				ss_adv_token,
				expected,
				ss_adv_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func getSnapshotRebalanceTest(t *testing.T) {
	tc := advancedSnapshotTestCases.GetAdvancedSnapshotRebalanceTestCase(t)

	expected, ok := tc.ExpectedResponse.(snapshotTester.ExpectedGetSnapshotRebalanceResponse)
	if !ok {
		t.Fatal("invalid ExpectedGetSnapshotRebalanceResponse")
	}

	snapshotTester.TestGetSnapshotRebalance(
		t,
		ss_adv_snapid,
		ss_adv_token,
		expected,
		ss_adv_testuser.User_id,
		tc.ExpectedStatusCode,
	)
}

func updateSnapshotValueOrderTest(t *testing.T) {
	t.Run("200 PUT", func(t2 *testing.T) {
		snapshotTester.TestUpdateSnapshotValueOrder(
			t2,
			ss_adv_snapid,
			snapshot.UpdateValueOrderPayload{
				Value_order: ss_adv_svids,
			},
			ss_adv_token,
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)
	})
}

func createSnapshotValueTest(t *testing.T) {
	var newSvId int
	t.Run("200 POST", func(t2 *testing.T) {
		newSvId = snapshotValueTester.TestCreateSnapshotValue(
			t2,
			snapshotvalue.CreateSnapshotValuePayload{
				Account_id:     ss_adv_accountids[0],
				Holding_id:     ss_adv_holdingids[8],
				Total:          100,
				Skip_rebalance: false,
			},
			ss_adv_token,
			ss_adv_snapid,
			ss_adv_testuser.User_id,
			fiber.StatusCreated,
		)
	})

	ss_adv_svids = append(ss_adv_svids, newSvId)

	// Since a value_order should have already been added, we expect the newly created snapshot to add to the existing value_order list
	t.Run("200 GET Verify", func(t2 *testing.T) {
		ss_adv_snapshot = snapshotTester.TestGetSnapshot(
			t,
			ss_adv_snapid,
			ss_adv_token,
			// Pass expected response with nil properties so we can skip the checks (only interested in the results of the GET call)
			snapshotTester.ExpectedGetSnapshotResponse{
				AccountIds:    nil,
				HoldingIds:    nil,
				Total:         0,
				WeightedErPct: 0,
			},
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)

		assert.Equal(t2, ss_adv_svids, ss_adv_snapshot.Value_order)
	})
}

var updateSnapshotTotal float64
var updateSnapshotExpenseRatio float64

func updateSnapshotValueTest(t *testing.T) {
	oldSvTotal := 10341.01 // See original create advanced snapshot test case
	newSvTotal := 650.99
	expectedNewSnapshotTotal := advancedSnapshotTestCases.AdvancedSnapshotTotal - oldSvTotal + newSvTotal

	// ER of holding 0 is 0
	erSum := advancedSnapshotTestCases.AdvancedSnapshotExpenseRatio * advancedSnapshotTestCases.AdvancedSnapshotTotal
	expectedNewErTotal := erSum / expectedNewSnapshotTotal

	updateSnapshotTotal = expectedNewSnapshotTotal
	updateSnapshotExpenseRatio = expectedNewErTotal

	// Round values
	expectedNewSnapshotTotal = math.Round(expectedNewSnapshotTotal*100) / 100
	expectedNewErTotal = math.Round(expectedNewErTotal*1000) / 1000

	t.Run("200 PUT", func(t2 *testing.T) {
		snapshotValueTester.TestUpdateSnapshotValue(
			t2,
			ss_adv_snapid,
			ss_adv_svids[0],
			snapshotvalue.UpdateSnapshotValuePayload{
				Account_id: ss_adv_accountids[0],
				Holding_id: ss_adv_holdingids[0],
				Total:      newSvTotal,
			},
			ss_adv_token,
			snapshotValueTester.ExpectedUpdateSnapshotValueResponse{
				Total: expectedNewSnapshotTotal,
				Er:    expectedNewErTotal,
			},
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)
	})

	t.Run("200 GET Verify", func(t2 *testing.T) {
		snapshotTester.TestGetSnapshot(
			t2,
			ss_adv_snapid,
			ss_adv_token,
			snapshotTester.ExpectedGetSnapshotResponse{
				AccountIds:    ss_adv_accountids,
				HoldingIds:    ss_adv_holdingids,
				Total:         expectedNewSnapshotTotal,
				WeightedErPct: expectedNewErTotal,
			},
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)
	})
}

func deleteSnapshotValueTest(t *testing.T) {
	deletedSvTotal := 650.99
	expectedNewSnapshotTotal := updateSnapshotTotal - deletedSvTotal

	// ER of holding 0 is 0
	erSum := updateSnapshotExpenseRatio * updateSnapshotTotal
	expectedNewErTotal := erSum / expectedNewSnapshotTotal

	// Round values
	expectedNewSnapshotTotal = math.Round(expectedNewSnapshotTotal*100) / 100
	expectedNewErTotal = math.Round(expectedNewErTotal*1000) / 1000

	t.Run("200 DEL", func(t2 *testing.T) {
		snapshotValueTester.TestDeleteSnapshotValue(
			t2,
			ss_adv_snapid,
			ss_adv_svids[0],
			ss_adv_token,
			snapshotValueTester.ExpectedDeleteSnapshotValueResponse{
				Total: expectedNewSnapshotTotal,
				Er:    expectedNewErTotal,
			},
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)
	})

	t.Run("200 GET Verify", func(t2 *testing.T) {
		ss_adv_snapshot = snapshotTester.TestGetSnapshot(
			t2,
			ss_adv_snapid,
			ss_adv_token,
			snapshotTester.ExpectedGetSnapshotResponse{
				AccountIds:    ss_adv_accountids,
				HoldingIds:    ss_adv_holdingids,
				Total:         expectedNewSnapshotTotal,
				WeightedErPct: expectedNewErTotal,
			},
			ss_adv_testuser.User_id,
			fiber.StatusOK,
		)
		assert.EqualExportedValues(t2, ss_adv_svids[1:], ss_adv_snapshot.Value_order)
	})
}

func deleteSnapshotTest(t *testing.T) {
	snapshotTester.TestDeleteSnapshot(
		t,
		ss_adv_snapid,
		ss_adv_token,
		ss_adv_testuser.User_id,
		fiber.StatusOK,
	)
}

func snapshotServiceCleaner(t *testing.T) {
	userTester.TestDeleteUser(t, "", ss_adv_token, ss_adv_testuser.User_id, fiber.StatusOK)
}
