package user

import (
	"testing"

	"github.com/MicroFish91/portfolio-instruments-api/api/services/user"
	"github.com/MicroFish91/portfolio-instruments-api/tests/integration/shared"
	"github.com/MicroFish91/portfolio-instruments-api/tests/utils"
	"github.com/gofiber/fiber/v3"
)

func UpdateSettingsTestCases(t *testing.T, userId int, email string) []shared.TestCase {
	tok401, tok403, err := utils.Generate40xTokens(userId, email)
	if err != nil {
		t.Fatal(err)
	}

	return []shared.TestCase{
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
