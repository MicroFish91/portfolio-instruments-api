package middleware

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
	"github.com/gofiber/fiber/v3"
)

func FallbackHandler(c fiber.Ctx, err error) error {
	return utils.SendError(c, fiber.StatusInternalServerError, err)
}
