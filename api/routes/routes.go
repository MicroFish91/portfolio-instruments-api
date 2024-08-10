package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/gofiber/fiber/v3"
)

func RegisterRoutes(
	app *fiber.App,
	authHandler types.AuthHandler,
	userHandler types.UserHandler,
	accountHandler types.AccountHandler,
	holdingHandler types.HoldingHandler,
	benchmarkHandler types.BenchmarkHandler,
	snapshotHandler types.SnapshotHandler,
	snapshotValueHandler types.SnapshotValueHandler,
) {
	app.Get("/ping", func(c fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{"data": "pong"})
	})

	routerV1 := app.Group("/api/v1")

	registerAuthRoutes(routerV1, authHandler)
	registerUserRoutes(routerV1, userHandler)
	RegisterBenchmarkRoutes(routerV1, benchmarkHandler)
	registerAccountRoutes(routerV1, accountHandler)
	registerHoldingRoutes(routerV1, holdingHandler)
	registerSnapshotRoutes(routerV1, snapshotHandler)
	registerSnapshotValueRoutes(routerV1, snapshotValueHandler)
}
