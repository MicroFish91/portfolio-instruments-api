package benchmark

import (
	"errors"
	"regexp"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *BenchmarkHandlerImpl) DeleteBenchmark(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	benchmarkParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(DeleteBenchmarkById)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid benchmark params from request"))
	}

	settings, err := h.userStore.GetSettings(c.Context(), userPayload.User_id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	if settings.Benchmark_id == benchmarkParams.Id {
		return utils.SendError(c, fiber.StatusConflict, errors.New("provided benchmark is actively saved as the target in user settings and cannot be deleted"))
	}

	benchmark, err := h.benchmarkStore.DeleteBenchmark(c.Context(), userPayload.User_id, benchmarkParams.Id)

	if err != nil {

		isStillInUse := regexp.MustCompile(`(?i)violates\sforeign\skey\sconstraint`).Match([]byte(err.Error()))
		if isStillInUse {

			benchmark, err = h.benchmarkStore.GetBenchmarkById(c.Context(), userPayload.User_id, benchmarkParams.Id)
			if err != nil {
				return utils.SendError(c, utils.StatusCodeFromError(err), err)
			}

			if !benchmark.Is_deprecated {
				benchmark, err = h.benchmarkStore.UpdateBenchmark(c.Context(), &types.Benchmark{
					Benchmark_id:     benchmarkParams.Id,
					Name:             benchmark.Name,
					Description:      benchmark.Description,
					Asset_allocation: benchmark.Asset_allocation,
					Std_dev_pct:      benchmark.Std_dev_pct,
					Real_return_pct:  benchmark.Real_return_pct,
					Drawdown_yrs:     benchmark.Drawdown_yrs,
					Is_deprecated:    true,
					User_id:          userPayload.User_id,
				})

				if err != nil {
					return utils.SendError(c, utils.StatusCodeFromError(err), err)
				}
			}

			return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
				"message":   "Resource is still being referenced by an existing snapshot and could not be deleted. Resource marked as deprecated instead and will be deleted automatically when no longer in use.",
				"benchmark": benchmark,
			})
		}

		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"message":   "Resource deleted successfully.",
		"benchmark": benchmark,
	})
}
