package benchmark

import (
	"context"
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func (h *BenchmarkHandlerImpl) GetBenchmarkById(c fiber.Ctx) error {
	userPayload, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
	if !ok {
		return utils.SendError(c, fiber.StatusUnauthorized, errors.New("unable to parse valid user from token"))
	}

	benchmarkParams, ok := c.Locals(constants.LOCALS_REQ_PARAMS).(GetBenchmarkByIdParams)
	if !ok {
		return utils.SendError(c, fiber.StatusBadRequest, errors.New("unable to parse valid account params from request"))
	}

	ctx, cancel := context.WithTimeout(c.Context(), time.Second*5)
	defer cancel()

	benchmark, err := h.store.GetBenchmarkById(ctx, userPayload.User_id, benchmarkParams.Id)
	if err != nil {
		return utils.SendError(c, utils.StatusCodeFromError(err), err)
	}

	return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"benchmark": benchmark})
}
