package routes

import (
	"github.com/MicroFish91/portfolio-instruments-api/api/types"
	"github.com/MicroFish91/portfolio-instruments-api/api/utils"
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
		return utils.SendJSON(c, fiber.StatusOK, fiber.Map{"message": "pong"})
	})

	app.All("/api/v1*", func(c fiber.Ctx) error {
		return utils.SendJSON(c, fiber.StatusGone, fiber.Map{"message": "The v1 API has been officially deprecated, please use the v2 API"})
	})

	routerV2 := app.Group("/api/v2")

	registerAuthRoutes(routerV2, authHandler)
	registerUserRoutes(routerV2, userHandler)
	registerBenchmarkRoutes(routerV2, benchmarkHandler)
	registerAccountRoutes(routerV2, accountHandler)
	registerHoldingRoutes(routerV2, holdingHandler)
	registerSnapshotRoutes(routerV2, snapshotHandler)
	registerSnapshotValueRoutes(routerV2, snapshotValueHandler)
}
