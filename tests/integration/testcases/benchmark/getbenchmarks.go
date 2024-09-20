package benchmark

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

type GetBenchmarksExpectedResponse struct {
	Benchmarks int
	Pagination types.PaginationMetadata
}

func GetBenchmarksTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
		// ---- 200 ----
		{
			Title:              "200",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 30,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 1",
			Route:              "/api/v1/benchmarks?current_page=1&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 20,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Page 2",
			Route:              "/api/v1/benchmarks?current_page=2&page_size=20",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 10,
				Pagination: types.PaginationMetadata{
					Current_page: 2,
					Page_size:    20,
					Total_items:  30,
				},
			},
		},
		{
			Title:              "200 Query Ids",
			Route:              "/api/v1/benchmarks?ids=1,11,15,29,40",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 4,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  4,
				},
			},
		},
		{
			Title:              "200 Query Name",
			Route:              "/api/v1/benchmarks?name=Classic Portfolio 1",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},
		{
			Title:              "200 Is Deprecated",
			Route:              "/api/v1/benchmarks?is_deprecated=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
		},
		{
			Title:              "200 Combination",
			Route:              "/api/v1/benchmarks?ids=1,11,15,29&is_deprecated=true",
			ExpectedStatusCode: fiber.StatusOK,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 1,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  1,
				},
			},
		},

		// ---- 401, 404 ----
		{
			Title:              "401",
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "404",
			Route:              "/api/v1/benchmarks?ids=40",
			ExpectedStatusCode: fiber.StatusNotFound,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 0,
				Pagination: types.PaginationMetadata{},
			},
		},

		// ---- 400 ----
		{
			Title:              "400 Ids",
			Route:              "/api/v1/benchmarks?ids=eleven",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Is Deprecated",
			Route:              "/api/v1/benchmarks?is_deprecated=1",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 0,
				Pagination: types.PaginationMetadata{},
			},
		},
		{
			Title:              "400 Combination",
			Route:              "/api/v1/benchmarks?is_deprecated=true&ids=true",
			ExpectedStatusCode: fiber.StatusBadRequest,
			ExpectedResponse: GetBenchmarksExpectedResponse{
				Benchmarks: 0,
				Pagination: types.PaginationMetadata{},
			},
		},
	}
}
