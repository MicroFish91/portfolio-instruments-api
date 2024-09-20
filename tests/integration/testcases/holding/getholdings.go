package holding

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/servicereqs/holding"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetHoldingsTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 200 ----
		{
			Title:              "200",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 30,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 1",
			Route:              "/api/v1/holdings?current_page=1&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 20,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 2",
			Route:              "/api/v1/holdings?current_page=2&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 10,
				Pagination: types.PaginationMetadata{
					Current_page: 2,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Ids",
			Route:              "/api/v1/holdings?ids=2,5,10,15,40",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 4,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  4,
				},
			},
		},
		{
			Title:              "200 Query Ticker",
			Route:              "/api/v1/holdings?ticker=T2",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},
		{
			Title:              "200 Query Asset Category (AC)",
			Route:              "/api/v1/holdings?asset_category=STB",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 5,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  5,
				},
			},
		},
		{
			Title:              "200 Query Asset Category (AC) 2",
			Route:              "/api/v1/holdings?asset_category=CASH",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Query Is Deprecated",
			Route:              "/api/v1/holdings?is_deprecated=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Query Ticker & AC",
			Route:              "/api/v1/holdings?ticker=T3&asset_category=DSCV",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},
		{
			Title:              "200 Query Maturation Remaining 1",
			Route:              "/api/v1/holdings?has_maturation_remaining=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 12,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  12,
				},
			},
		},
		{
			Title:              "200 Query Maturation Remaining 2",
			Route:              "/api/v1/holdings?has_maturation_remaining=false",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},
		{
			Title:              "200 Query AC & Maturation Remaining 1",
			Route:              "/api/v1/holdings?asset_category=LTB&has_maturation_remaining=false",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},
		{
			Title:              "200 Query AC & Maturation Remaining 2",
			Route:              "/api/v1/holdings?asset_category=LTB&has_maturation_remaining=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},

		// ---- 401, 404 ----
		{
			Title:              "401",
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "404",
			Route:              "/api/v1/holdings?ids=40",
			ExpectedStatusCode: fiber.StatusNotFound,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "404 Query Ticker & AC",
			Route:              "/api/v1/holdings?ticker=T3&asset_category=TSM",
			ExpectedStatusCode: fiber.StatusNotFound,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 0,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  0,
				},
			},
		},
		{
			Title:              "404 Query Ticker & AC & IsDep",
			Route:              "/api/v1/holdings?ticker=T3&asset_category=DSCV&is_deprecated=true",
			ExpectedStatusCode: fiber.StatusNotFound,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings: 0,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  0,
				},
			},
		},

		// ---- 400 ----
		{
			Title:              "400 Bad Ids",
			Route:              "/api/v1/holdings?ids=2.0,3.0",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Bad Asset Category 1",
			Route:              "/api/v1/holdings?asset_category=10",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Bad Asset Category 2",
			Route:              "/api/v1/holdings?asset_category=OIL",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Bad Is Deprecated",
			Route:              "/api/v1/holdings?is_deprecated=maybe",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Bad Maturation Remaining",
			Route:              "/api/v1/holdings?has_maturation_remaining=06/11/2029",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: holding.GetHoldingsExpectedResponse{
				Holdings:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
	}
}
