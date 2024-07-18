package benchmark

import (
	"errors"
	"fmt"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *BenchmarkHandlerImpl) CreateBenchmark(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	benchmarkPayload, ok := c.Locals(constants.LOCALS_REQ_BODY).(CreateBenchmarkPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid benchmark payload from request"))
	}

	// Ensure benchmark name is unique per user
	existingBenchmark, _ := h.store.GetBenchmarkByName(c.Context(), benchmarkPayload.Name, userPayload.User_id)
	if existingBenchmark != nil {
		return utils.SendError(c, fiber.StatusConflict, fmt.Errorf(`user already has existing benchmark with name "%s"`, existingBenchmark.Name))
	}

	benchmark, err := h.store.CreateBenchmark(
		c.Context(),
		&types.Benchmark{
			Name:             benchmarkPayload.Name,
			Description:      benchmarkPayload.Description,
			Asset_allocation: benchmarkPayload.Asset_allocation,
			Std_dev_pct:      benchmarkPayload.Std_dev_pct,
			Real_return_pct:  benchmarkPayload.Real_return_pct,
			Drawdown_yrs:     benchmarkPayload.Drawdown_yrs,
			Is_deprecated:    benchmarkPayload.Is_deprecated,
			User_id:          userPayload.User_id,
		},
	)

	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}
	return utils.SendJSON(c, fiber.StatusCreated, fiber.Map{"benchmark": benchmark})
}
