package usercases

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/testcase"
	"github.com/gofiber/fiber/v3"
)

func UpdateSettingsTestCases(t *testing.T, userId int) []testcase.PutTestCase {
	tok401, err := auth.GenerateSignedJwt(userId, "test_user@gmail.com", "Default")
	if err != nil {
		t.Fatal(err)
	}
	tok401 = tok401[1:]

	tok403, err := auth.GenerateSignedJwt(100, "fake_user_100@gmail.com", "Default")
	if err != nil {
		t.Fatal(err)
	}

	return []testcase.PutTestCase{
		{
			Title: "200",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 5,
			},
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusOK,
		},
		{
			Title: "401",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 5,
			},
			ParameterId:        userId,
			ReplacementToken:   tok401,
			ExpectedStatusCode: fiber.StatusUnauthorized,
		},
		{
			Title: "403 Token Id",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 5,
			},
			ParameterId:        userId,
			ReplacementToken:   tok403,
			ExpectedStatusCode: fiber.StatusForbidden,
		},
		{
			Title: "403 Param Id",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 5,
			},
			Route:              "/api/v1/users/100/settings",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusForbidden,
		},
		{
			Title: "409",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 5,
				Benchmark_id:   1,
			},
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusConflict,
		},

		// ---- 400 ----

		// Params
		{
			Title: "400 Params String Id",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 10,
				Benchmark_id:   1,
			},
			Route:              "/api/v1/users/test/settings",
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Params Float Id",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 10,
				Benchmark_id:   1,
			},
			Route:              "/api/v1/users/1.0/settings",
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
		{
			Title: "400 Params Object Id",
			Payload: user.UpdateSettingsPayload{
				Reb_thresh_pct: 10,
				Benchmark_id:   1,
			},
			Route:              "/api/v1/users/{id:1}/settings",
			ExpectedStatusCode: fiber.StatusBadRequest,
		},

		// Payload
		{
			Title: "400 Payload",
			Payload: map[string]any{
				"Reb_thresh_pct": "five",
				"Benchmark_id":   true,
			},
			ParameterId:        userId,
			ExpectedStatusCode: fiber.StatusBadRequest,
		},
	}
}
