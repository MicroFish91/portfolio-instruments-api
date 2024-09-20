package account

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

type GetAccountsExpectedResponse struct {
	Accounts   int
	Pagination types.PaginationMetadata
}

func GetAccountsTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 200 ----
		{
			Title:              "200",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 30,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 1",
			Route:              "/api/v1/accounts?current_page=1&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 20,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 2",
			Route:              "/api/v1/accounts?current_page=2&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 10,
				Pagination: types.PaginationMetadata{
					Current_page: 2,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Ids",
			Route:              "/api/v1/accounts?ids=1,11,15,31",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Query Ids",
			Route:              "/api/v1/accounts?ids=1,11,15,31",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Query Tax Shelter 1",
			Route:              "/api/v1/accounts?tax_shelter=TAXABLE",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 9,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  9,
				},
			},
		},
		{
			Title:              "200 Query Tax Shelter 2",
			Route:              "/api/v1/accounts?tax_shelter=ROTH",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 6,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  6,
				},
			},
		},
		{
			Title:              "200 Query Institution 1",
			Route:              "/api/v1/accounts?institution=vanguard",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 9,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  9,
				},
			},
		},
		{
			Title:              "200 Query Institution 2",
			Route:              "/api/v1/accounts?institution=fidelity",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 9,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  9,
				},
			},
		},
		{
			Title:              "200 Query Is Deprecated",
			Route:              "/api/v1/accounts?is_deprecated=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Query Shelter + Inst",
			Route:              "/api/v1/accounts?institution=fidelity&tax_shelter=ROTH",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts: 3,
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
			Route:              "/api/v1/accounts",
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "404",
			Route:              "/api/v1/accounts?ids=32",
			ExpectedStatusCode: fiber.StatusNotFound,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},

		// --- 400 ---
		{
			Title:              "400 Fake Query Param",
			Route:              "/api/v1/accounts?fake_query=true",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Tax Shelter 1",
			Route:              "/api/v1/accounts?tax_shelter=TRAD",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Tax Shelter 2",
			Route:              "/api/v1/accounts?tax_shelter=5",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Is Deprecated",
			Route:              "/api/v1/accounts?is_deprecated=test",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetAccountsExpectedResponse{
				Accounts:   0,
				Pagination: types.PaginationMetadata{},
			},
		},
	}
}
