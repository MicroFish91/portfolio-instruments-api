package integration

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/holding"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	holdingTestCases "github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases/holding"
	holdingTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/holding"
	userTester "github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

var (
	hs_token    string
	hs_testuser types.User

	holdId    int
	hs_eridx  int
	hs_depidx int
	hs_incidx int
	hs_mfidx  int
)

func TestHoldingService(t *testing.T) {
	t.Parallel()

	t.Run("Setup", holdingServiceSetup)
	t.Run("POST://api/v1/holdings", createHoldingTests)
	t.Run("GET://api/v1/holdings", getHoldingsTests)
	t.Run("GET://api/v1/holdings/:id", getHoldingTests)
	t.Run("PUT://api/v1/holdings/:id", updateHoldingTests)
	t.Run("DEL://api/v1/holdings/:id", deleteHoldingTests)
	t.Run("Cleanup", holdingServiceCleanup)
}

func holdingServiceSetup(t *testing.T) {
	hs_testuser, hs_token = newUserSetup(t)
}

func createHoldingTests(t *testing.T) {
	for _, tc := range holdingTestCases.GetCreateHoldingTestCases(t, hs_testuser.User_id, hs_token) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestCreateHolding(
				t2,
				tc.Payload,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func getHoldingsTests(t *testing.T) {
	// Get holdings test setup
	t.Run("Setup", func(t2 *testing.T) {
		for i := 0; i < 25; i += 1 {
			var mockAsset utils.MockAsset
			if i%2 == 0 {
				mockAsset = utils.GetRotatingFixedIncome(&hs_incidx)
			} else {
				mockAsset = utils.GetRotatingMutualFund(&hs_mfidx)
			}

			holdingTester.TestCreateHolding(
				t2,
				holding.CreateHoldingPayload{
					Name:              fmt.Sprintf("Hold%d", i),
					Ticker:            mockAsset.Ticker,
					Asset_category:    mockAsset.Asset_category,
					Expense_ratio_pct: utils.GetRotatingExpenseRatio(&hs_eridx),
					Maturation_date:   mockAsset.Maturation_date,
					Interest_rate_pct: mockAsset.Interest_rate_pct,
					Is_deprecated:     utils.GetRotatingDeprecation(&hs_depidx),
				},
				hs_token,
				hs_testuser.User_id,
				fiber.StatusCreated,
			)
		}

		// Create asset with an expired maturation date
		holdId = holdingTester.TestCreateHolding(
			t2,
			holding.CreateHoldingPayload{
				Name:              "Hold26",
				Asset_category:    "LTB",
				Maturation_date:   "01/01/1990",
				Interest_rate_pct: 5.8,
			},
			hs_token,
			hs_testuser.User_id,
			fiber.StatusCreated,
		)
	})

	// Get holdings tests
	for _, tc := range holdingTestCases.GetHoldingsTestsCases(t, hs_testuser.User_id, hs_token) {
		t.Run(tc.Title, func(t2 *testing.T) {
			response, ok := tc.ExpectedResponse.(holdingTestCases.GetHoldingsExpectedResponse)
			if !ok {
				t.Fatal("invalid GetHoldingsExpectedResponse")
			}

			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestGetHoldings(
				t2,
				tc.Route,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
				response,
			)
		})
	}
}

func getHoldingTests(t *testing.T) {
	for _, tc := range holdingTestCases.GetHoldingTestCases(t, holdId, hs_testuser.User_id, hs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestGetHolding(
				t2,
				tc.ParameterId,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func updateHoldingTests(t *testing.T) {
	for _, tc := range holdingTestCases.GetUpdateHoldingTestCases(t, holdId, hs_testuser.User_id, hs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestUpdateHolding(
				t2,
				tc.ParameterId,
				tc.Payload,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func deleteHoldingTests(t *testing.T) {
	for _, tc := range holdingTestCases.GetDeleteHoldingTestCases(t, holdId, hs_testuser.User_id, hs_testuser.Email) {
		t.Run(tc.Title, func(t2 *testing.T) {
			tok := hs_token
			if tc.ReplacementToken != "" {
				tok = tc.ReplacementToken
			}

			holdingTester.TestDeleteHolding(
				t2,
				tc.ParameterId,
				tok,
				hs_testuser.User_id,
				tc.ExpectedStatusCode,
			)
		})
	}
}

func holdingServiceCleanup(t *testing.T) {
	userTester.TestDeleteUser(t, "", hs_token, hs_testuser.User_id, fiber.StatusOK)
}
