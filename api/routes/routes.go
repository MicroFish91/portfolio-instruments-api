package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(app *fiber.App, userHandler types.UserHandler, accountHandler types.AccountHandler) {

	app.Get("/ping", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{"data": "pong"})
	})

	v1 := app.Group("/api/v1")
	registerAuthRoutes(v1, userHandler)
	registerAccountRoutes(v1, accountHandler)
}
