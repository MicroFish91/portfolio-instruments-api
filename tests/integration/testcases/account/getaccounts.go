package account

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/gofiber/fiber/v3"
)

type GetAccountsExpectedResponse struct {
	Accounts   int
	Pagination types.PaginationMetadata
}

// type GetAccountsQuery struct {
// 	Ids           []int            `json:"ids"`
// 	Tax_shelter   types.TaxShelter `json:"tax_shelter"`
// 	Institution   string           `json:"institution"`
// 	Is_deprecated string           `json:"is_deprecated"`

// 	types.PaginationQuery
// }

func GetAccountsTestCases() []shared.GetTestCase {
	return []shared.GetTestCase{
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
	}
}
