package middleware

import (
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func AddRateLimiter(limit int, expiration time.Duration) func(fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        limit,
		Expiration: expiration,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			return utils.SendError(c, fiber.StatusTooManyRequests, errors.New("short term request limit exceeded"))
		},
	})
}
