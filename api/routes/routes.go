package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(
	app *fiber.App,
	userHandler types.UserHandler,
	accountHandler types.AccountHandler,
	holdingHandler types.HoldingHandler,
	benchmarkHandler types.BenchmarkHandler,
	snapshotHandler types.SnapshotHandler,
) {
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{"data": "pong"})
	})

	v1 := app.Group("/api/v1")
	registerUserRoutes(v1, userHandler)
	RegisterBenchmarkRoutes(v1, benchmarkHandler)
	registerAccountRoutes(v1, accountHandler)
	registerHoldingRoutes(v1, holdingHandler)
	registerSnapshotRoutes(v1, snapshotHandler)
}
