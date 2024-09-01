package benchmark

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *BenchmarkHandlerImpl) UpdateBenchmark(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	benchmarkPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(UpdateBenchmarkPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid benchmark payload from request"))
	}

	benchmarkParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(UpdateBenchmarkById)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid benchmark params from request"))
	}

	// Ensure no duplicate benchmarks per user
	b, _ := h.benchmarkStore.GetBenchmarkByName(c.Context(), benchmarkPayload.Name, userPayload.User_id)
	if b.Benchmark_id != 0 && b.Benchmark_id != benchmarkParams.Id {
		return utils.SendError(c, fiber.StatusConflict, errors.New("benchmark with the given name already exists for the user"))
	}

	benchmark, err := h.benchmarkStore.UpdateBenchmark(c.Context(), types.Benchmark{
		Benchmark_id:     benchmarkParams.Id,
		Name:             benchmarkPayload.Name,
		Description:      benchmarkPayload.Description,
		Asset_allocation: benchmarkPayload.Asset_allocation,
		Std_dev_pct:      benchmarkPayload.Std_dev_pct,
		Real_return_pct:  benchmarkPayload.Real_return_pct,
		Drawdown_yrs:     benchmarkPayload.Drawdown_yrs,
		Is_deprecated:    benchmarkPayload.Is_deprecated,
		User_id:          userPayload.User_id,
	})

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"benchmark": benchmark})
}
