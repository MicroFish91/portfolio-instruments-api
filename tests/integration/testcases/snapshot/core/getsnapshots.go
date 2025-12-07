package core

import (
	"fmt"
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	routeTester "github.com/MicroFish91/portfolio-instruments-api/tests/integration/routetester/snapshot"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/testcases"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func GetSnapshotsTestCases(t *testing.T, snapshotId, userId int, email string) []testcases.TestCase {
	tok401, _, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []testcases.TestCase{
		// ---- 200 ----
		{
			Title: "200",
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 31,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  31,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Pagination 1",
			Route: "/api/v2/snapshots?page_size=20",
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 20,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    20,
					Total_items:  31,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Pagination 2",
			Route: "/api/v2/snapshots?current_page=2&page_size=20",
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 11,
				Pagination: types.PaginationMetadata{
					Current_page: 2,
					Page_size:    20,
					Total_items:  31,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Ids",
			Route: fmt.Sprintf("/api/v2/snapshots?snap_ids=%d,%d,%d,%d", snapshotId, snapshotId+2, snapshotId+5, snapshotId+100),
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 3,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  3,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Ascending",
			Route: "/api/v2/snapshots?order_date_by=ASC",
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 31,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  31,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "200 Descending",
			Route: "/api/v2/snapshots?order_date_by=DESC",
			ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
				Snapshots: 31,
				Pagination: types.PaginationMetadata{
					Current_page: 1,
					Page_size:    50,
					Total_items:  31,
				},
			},
			ExpectedStatusCode: fiber.StatusOK,
		},
		// Generally works, but can be flaky on certain days due to having to match up with variable dates... probably would have been a better idea to use static test dates rather than relative ones
		// {
		// 	Title: "200 Date Lower",
		// 	Route: fmt.Sprintf("/api/v2/snapshots?snap_date_lower=%s", utils.Calc_target_date(-2, 0)),
		// 	ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
		// 		Snapshots: 12,
		// 		Pagination: types.PaginationMetadata{
		// 			Current_page: 1,
		// 			Page_size:    50,
		// 			Total_items:  12,
		// 		},
		// 	},
		// 	ExpectedStatusCode: fiber.StatusOK,
		// },
		// {
		// 	Title: "200 Date Upper",
		// 	Route: fmt.Sprintf("/api/v2/snapshots?snap_date_upper=%s", utils.Calc_target_date(-2, 0)),
		// 	ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
		// 		Snapshots: 18,
		// 		Pagination: types.PaginationMetadata{
		// 			Current_page: 1,
		// 			Page_size:    50,
		// 			Total_items:  18,
		// 		},
		// 	},
		// 	ExpectedStatusCode: fiber.StatusOK,
		// },
		// {
		// 	Title: "200 Date Lower Upper",
		// 	Route: fmt.Sprintf("/api/v2/snapshots?snap_date_lower=%s&snap_date_upper=%s", utils.Calc_target_date(-3, 0), utils.Calc_target_date(-2, 0)),
		// 	ExpectedResponse: routeTester.ExpectedGetSnapshotsResponse{
		// 		Snapshots: 13,
		// 		Pagination: types.PaginationMetadata{
		// 			Current_page: 1,
		// 			Page_size:    50,
		// 			Total_items:  13,
		// 		},
		// 	},
		// 	ExpectedStatusCode: fiber.StatusOK,
		// },

		// ---- 401 ----
		{
			Title:              "401",
			ReplacementToken:   tok401,
			ExpectedResponse:   routeTester.ExpectedGetSnapshotsResponse{},
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},

		// ---- 400 ----
		{
			Title:              "400 Ids",
			Route:              "/api/v2/snapshots?ids=1.0,2.0",
			ExpectedResponse:   routeTester.ExpectedGetSnapshotsResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Date Lower",
			Route:              "/api/v2/snapshots?snap_date_lower=\"2022/10/02\"",
			ExpectedResponse:   routeTester.ExpectedGetSnapshotsResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Date Upper",
			Route:              "/api/v2/snapshots?snap_date_upper=5",
			ExpectedResponse:   routeTester.ExpectedGetSnapshotsResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title:              "400 Order By",
			Route:              "/api/v2/snapshots?order_date_by=true",
			ExpectedResponse:   routeTester.ExpectedGetSnapshotsResponse{},
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
