package middleware

import (
	"log/slog"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func AddLocalsContextLogger(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		c.Locals(constants.LOCALS_LOGGER, logger)
		return c.Next()
	}
}

func AddIncomingTrafficLogger(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		id := uuid.New().String()[:8]
		c.Locals(constants.LOCALS_REQ_ID, id)
		c.Locals(constants.LOCALS_REQ_START, time.Now())

		logger.Info(
			"Inbound traffic: ",
			"Id", id,
			"Method", c.Method(),
			"Path", c.Path(),
			"IP", c.IP(),
		)

		return c.Next()
	}
}
