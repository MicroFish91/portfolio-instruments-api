package middleware

import (
	"errors"
	"time"

	"github.com/MicroFish91/portfolio-instruments-api/api/constants"
	"github.com/MicroFish91/portfolio-instruments-api/api/services/auth"
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
			return utils.SendError(c, fiber.StatusTooManyRequests, errors.New("request limit exceeded"))
		},
	})
}

func AddUnauthorizedRateLimiter() func(fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:        30,
		Expiration: 60 * time.Minute,
		Next: func(c fiber.Ctx) bool {
			authUser, ok := c.Locals(constants.LOCALS_REQ_USER).(auth.AuthUserPayload)
			return ok && authUser.Email != "" && authUser.User_id != 0
		},
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			return utils.SendError(c, fiber.StatusTooManyRequests, errors.New("request limit exceeded"))
		},
	})
}
