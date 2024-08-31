package benchmark

import (
	"errors"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *BenchmarkHandlerImpl) GetBenchmarks(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	queryPayload, ok := c.Locals(constants.LOCALS_REQ_QUERY).(GetBenchmarksQuery)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid query params from request"))
	}

	benchmarks, pagination, err := h.benchmarkStore.GetBenchmarks(
		c.Context(),
		userPayload.User_id,
		types.GetBenchmarksStoreOptions{
			Benchmark_ids: queryPayload.Ids,
			Name:          queryPayload.Name,
			Is_deprecated: queryPayload.Is_deprecated,
			Current_page:  queryPayload.Current_page,
			Page_size:     queryPayload.Page_size,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{
		"benchmarks": benchmarks,
		"pagination": pagination,
	})
}
